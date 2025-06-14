# === YASM: Fish shell keybinding setup ===

# Define the function triggered by Ctrl-G
function __yasm_run_with_injection
    set -l output (yasm run 2>&1)
    set -l exit_code $status

    if test $exit_code -eq 0
        commandline -r -- "$output"
    else
        # Print all lines except the last with newline
        for i in (seq (math (count $output) - 1))
            echo $output[$i] >&2
        end

        # Print the last line without a newline because
        # `commandline -f cancel execute` add extra newline
        if test (count $output) -gt 0
            printf '%s' $output[-1] >&2
        end

        # This is to cancel the wait after printing the error
        commandline -f cancel execute
    end
end

# Function to (re)bind the key
function __yasm_setup_bindings --on-event fish_prompt
    # Remove any existing bindings to avoid duplicates
    for mode in insert default
        bind -M $mode -e ctrl-g 2>/dev/null
        bind -M $mode -e \cg 2>/dev/null
    end

    # Bind Ctrl-G using both readable and legacy forms for max compatibility
    bind -M insert ctrl-g __yasm_run_with_injection
    bind -M insert \cg __yasm_run_with_injection
    bind -M default ctrl-g __yasm_run_with_injection
    bind -M default \cg __yasm_run_with_injection
end

# Trigger the binding setup immediately
emit fish_prompt

# === YASM: Fish shell keybinding setup ===
