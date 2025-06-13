function __yasm_run_injector_widget() {
  local output
  output="$(yasm run 2>&1)"
  local exit_code=$?

  # If exited successfully, inject the output into the command line.
  # else just print the output.
  if [[ $exit_code -eq 0 ]]; then
    LBUFFER="$output"
  else
    print -r -- "$output"
  fi

  zle reset-prompt
}

zle -N __yasm_run_injector_widget
bindkey '^Y' __yasm_run_injector_widget

