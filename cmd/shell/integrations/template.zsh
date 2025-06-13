function _yasm_widget() {
  local output
  output="$(go run . run 2>&1)"
  local exit_code=$?

  # Check for ::inject::
  if [[ $exit_code -eq 0 && "$output" == *"{{INJECT_PREFIX}}"* ]]; then
    local injected_line
    injected_line=$(echo "$output" | grep '^{{INJECT_PREFIX}}' | sed 's/^{{INJECT_PREFIX}}//')
    LBUFFER="$injected_line"
  else
    print -r -- "$output"
  fi

  zle reset-prompt
}


zle -N _yasm_widget
bindkey '^Y' _yasm_widget

