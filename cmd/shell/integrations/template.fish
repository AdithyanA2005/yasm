# function yasm
#     set tmpfile (mktemp)
#
#     # Run command, tee output and capture exit code immediately
#     command go run . $argv | tee $tmpfile
#     set exit_code $status
#
#     if grep -q "::inject::" $tmpfile
#         set cmd (grep "::inject::" $tmpfile | sed 's/.*::inject:://')
#         commandline --replace $cmd
#         commandline --function repaint
#     end
#
#     rm -f $tmpfile
#     return $exit_code
# end

# function yasm
#     set tmpfile (mktemp)
#
#     # Use command substitution to preserve TTY for interactive programs like nvim or fzf
#     # Duplicate stdout using `tee`, but preserve terminal behavior
#     command go run . $argv ^ | tee $tmpfile
#     set exit_code $status
#
#     if grep -q "::inject::" $tmpfile
#         set cmd (grep "::inject::" $tmpfile | sed 's/.*::inject:://')
#         commandline --replace $cmd
#         commandline --function repaint
#     end
#
#     rm -f $tmpfile
#     return $exit_code
# end

# function yasm
#     set tmpfile (mktemp)
#
#     # Use 'script' to capture output while preserving TTY
#     # '-q' = quiet, '-c' = run this command, and output goes to tmpfile
#     script -q -c "go run . $argv" $tmpfile
#     set exit_code $status
#
#     # Show output from tmpfile (as the command normally would)
#     # cat $tmpfile
#
#     # Check and inject
#     if grep -q "::inject::" $tmpfile
#         set cmd (grep "::inject::" $tmpfile | sed 's/.*::inject:://')
#         commandline --replace $cmd
#         commandline --function repaint
#     end
#
#     rm -f $tmpfile
#     return $exit_code
# end
#
# -q – quiet mode (don't echo start message).
#
# -e – don’t print the "Script done" message.
#
# -f – flush output as it’s written (useful for real-time behavior).
#
# -c – the command to run.

function yasm
    set tmpfile (mktemp -t yasm_output.XXXXXX)

    # Run the CLI tool inside a real TTY, and suppress the "Script done" message
    script -qefc "go run . $argv" $tmpfile
    set exit_code $status
    # Repaint regardless (restores visibility)

    commandline --function repaint

    if grep -q "::inject::" $tmpfile
        set cmd (grep "::inject::" $tmpfile | sed 's/.*::inject:://')
        echo "CMD: $cmd"
        commandline --replace $cmd
        commandline --function repaint
    end

    # Just in case terminal is left in a bad state
    if not set -q __fish_initialized
        echo
        echo "hell no"
        echo
        stty sane
        commandline --function repaint
    end

    echo "TMP START"
    cat $tmpfile
    echo "TMP END"
    rm -f $tmpfile
    return $exit_code
end
