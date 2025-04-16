#!/usr/bin/env bash

# Function: check_installed
# Purpose : Check if the given list of commands are installed
# Args    : List of command names to check
# Returns : 0 if all dependencies are found (returns it as space-separated list), 1 otherwise

check_installed() {
  local missing=()
  for dep in "$@"; do
    if ! command -v "$dep" &>/dev/null; then
      missing+=("$dep") # Add missing dependency to the list
    fi
  done

  # Return the missing dependencies as a string
  if [ ${#missing[@]} -gt 0 ]; then
    echo "${missing[@]}" # Return the missing dependencies as a space-separated list
    return 1             # Return 1 to indicate missing dependencies
  fi

  return 0 # Return 0 if no dependencies are missing
}
