#!/usr/bin/env bash

function fzf_list_scripts() {
  local choice
  local script_name

  # Ensure that the META_DIR is not empty
  shopt -s nullglob
  local toml_files=("$META_DIR"/*.toml)
  shopt -u nullglob

  if [[ ${#toml_files[@]} -eq 0 ]]; then
    echo "No scripts found." >&2
    return 1
  fi

  # Generate list of scripts with descriptions
  choice=$(
    for file in "$META_DIR"/*.toml; do
      name=$(basename "$file" .toml)
      desc=$(grep -m1 description "$file" | cut -d'"' -f2)
      printf "%-20s - %s\n" "$name" "$desc"
    done | fzf --prompt=" Select script: " "${FZF_OPTS[@]}"
  )

  # Extract the script name from the selection
  script_name=$(echo "$choice" | awk '{print $1}') # We are only using the fist part before space as filename is expected to be one word

  # Returning the script name
  echo $script_name
}

function fzf_menu() {
  local choice
  local options=(
    "1. Run Script"
    "2. Create Script"
    "3. Update Script"
    "4. Rename Script"
    "5. Delete Script"
    "6. Show Help"
  )

  # Display options using fzf with newline-separated entries
  choice=$(printf "%s\n" "${options[@]}" |
    fzf --prompt=" Select action: " "${FZF_OPTS[@]}")

  # Handle the user's selection
  case "$choice" in
  "${options[0]}") # Run Script
    script_name=$(fzf_list_scripts)
    [ -z "$script_name" ] && echo "No script selected" || run_script "$script_name"
    ;;
  "${options[1]}") # Create Script
    create_script
    ;;
  "${options[2]}") # Update Script
    script_name=$(fzf_list_scripts)
    [ -z "$script_name" ] && echo "No script selected" || update_script "$script_name"
    ;;
  "${options[3]}") # Rename Script
    script_name=$(fzf_list_scripts)
    [ -z "$script_name" ] && echo "No script selected" || rename_script "$script_name"
    ;;
  "${options[4]}") # Delete Script
    script_name=$(fzf_list_scripts)
    [ -z "$script_name" ] && echo "No script selected" || delete_script "$script_name"
    ;;
  "${options[5]}") # Delete Script
    show_help
    ;;
  *)
    echo "No action selected."
    ;;
  esac
}
