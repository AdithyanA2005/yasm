# === YASM: ZSH shell keybinding setup ===

function __yasm_run_with_injection() {
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

zle -N __yasm_run_with_injection
bindkey '^G' __yasm_run_with_injection

# === YASM: ZSH shell keybinding setup ===
