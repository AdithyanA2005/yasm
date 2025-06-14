# === YASM: ZSH shell keybinding setup ===

__yasm_run_with_injection() {
  # Capture output of yasm run
  local output
  output="$(yasm run 2>&1)"
  local exit_code=$?

  # If success, inject output into command line
  if [[ $exit_code -eq 0 ]]; then
    READLINE_LINE="$output"
    READLINE_POINT=${#READLINE_LINE}
  else
    echo "$output"
  fi
}

bind -x '"\C-g":__yasm_run_with_injection'

# === YASM: ZSH shell keybinding setup ===
