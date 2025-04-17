#!/usr/bin/env bash

function create_script() {
  # Prompt user for a new script name
  echo -n "󰛿 Enter script name: "
  read -r name

  # Exit if no name was not entered
  [[ -z "$name" ]] && echo " No name entered. Script creation cancelled." && return

  # Exit if name is not valid
  [[ ! "$name" =~ ^[a-zA-Z0-9_-]+$ ]] && echo " Invalid name. Only alphanumeric, hyphens (-), and underscores (_) are allowed." && return

  local script_path="$SCRIPTS_DIR/$name.sh"
  local meta_path="$META_DIR/$name.toml"

  # Exit if script or metadata already exists
  [[ -f "$script_path" ]] && echo " Script for '$name' already exists." && return
  [[ -f "$meta_path" ]] && echo " Metadata for '$name' already exists." && return

  # Create the script and metadata files
  generate_script_template "$name" >"$script_path"
  generate_meta_template "$name" >"$meta_path"

  # Make the script executable
  chmod +x "$script_path"

  # Confirmation message
  echo " Script '$name' created!"
  echo "  - Script: $script_path"
  echo "  - Metadata: $meta_path"

  # Prompt user to open files in editor
  read -rp " Open files in editor($EDITOR)? [Y/n]: " open_choice

  # Convert to lowercase
  open_choice=$(echo "$open_choice" | tr '[:upper:]' '[:lower:]')

  # If user said yes open it in $EDITOR or nvim
  if [[ -z "$open_choice" || "$open_choice" == "y" || "$open_choice" == "yes" ]]; then
    ${EDITOR:-nano} "$script_path" "$meta_path"
  fi
}

function rename_script() {
  local old_name="$1"
  local old_script="$SCRIPTS_DIR/$old_name.sh"
  local old_meta="$META_DIR/$old_name.toml"

  # Exit if the old script or metadata file doesn't exist
  [[ ! -f "$old_script" ]] && echo " Script for '$old_name' not found." && return
  [[ ! -f "$old_meta" ]] && echo " Metadata for '$old_name' not found." && return

  # Take the new name for  script from user
  echo -n "󰚰 Enter the new name: "
  read -r new_name

  # Exit if a new name wasn't entered
  [[ -z "$new_name" ]] && echo " No new name entered. Rename cancelled." && return

  # Exit if the new name is not valid
  [[ ! "$new_name" =~ ^[a-zA-Z0-9_-]+$ ]] && echo " Invalid name. Only alphanumeric, hyphens (-), and underscores (_) are allowed." && return

  local new_script="$SCRIPTS_DIR/$new_name.sh"
  local new_meta="$META_DIR/$new_name.toml"

  # Exit if the new script or metadata file already exists
  [[ -e "$new_script" ]] && echo " Script for $new_name already exists." && return
  [[ -e "$new_meta" ]] && echo " Metadata for $new_name already exists." && return

  # Rename files
  mv "$old_script" "$new_script"
  mv "$old_meta" "$new_meta"

  # Update metadata
  sed -i '' "s/^name = \".*\"/name = \"$new_name\"/" "$new_meta"
  sed -i '' "s/^usage = \".*\"/usage = \"$new_name.sh <arguments>\"/" "$new_meta"

  # Show confirmation message
  echo " Script renamed from '$old_name' to '$new_name'."
}

function update_script() {
  local name="$1"
  local script_path="$SCRIPTS_DIR/$name.sh"
  local meta_path="$META_DIR/$name.toml"

  # Exit if the script or metadata file doesn't exist
  [[ ! -f "$script_path" ]] && echo " Script for '$name' not found." && return
  [[ ! -f "$meta_path" ]] && echo " Metadata for '$name' not found." && return

  # Open the script and metadata files in the editor
  ${EDITOR:-nano} "$script_path" "$meta_path"
}

function delete_script() {
  local name="$1"
  local script_path="$SCRIPTS_DIR/$name.sh"
  local meta_path="$META_DIR/$name.toml"

  # Exit if the script or metadata for file doesn't exist
  [[ ! -f "$script_path" ]] && echo " Script for '$name' not found." && return
  [[ ! -f "$meta_path" ]] && echo " Metadata for '$name' not found." && return

  # Confirm deletion
  echo -n " Are you sure you want to delete '$name'? (y/N): "
  read -r confirm

  # If user said yes delete it
  if [[ "$confirm" == "y" ]]; then
    rm -f "$script_path" "$meta_path"                           # Delete the script and metadata
    echo " Script '$name' and its metadata have been deleted." # Show confirmation message
  else
    echo " Deletion cancelled."
  fi
}

function run_script() {
  local name="$1"
  local script_path="$SCRIPTS_DIR/$name.sh"
  local meta_path="$META_DIR/$name.toml"

  # Exit if the script is not executable or doesn't exist
  [[ ! -x "$script_path" ]] && echo " Script for '$name' not found or not executable at: $script_path" && return

  # Exit if the metadata file doesn't exist
  [[ ! -f "$meta_path" ]] && echo " Metadata for '$name' not found at: $meta_path" && return

  # Check for dependencies in the metadata file
  if [[ -f "$meta_path" ]]; then
    # Extract dependencies from the metadata file
    local deps
    deps=$(grep '^dependencies' "$meta_path" | cut -d'[' -f2 | cut -d']' -f1 | tr ',' '\n' | tr -d '" ')

    if [[ -n "$deps" ]]; then
      # Capture the missing dependencies
      missing_deps=$(check_installed $deps)

      # If there are missing dependencies, return and don't run the script
      if [[ -n "$missing_deps" ]]; then
        echo " Missing dependencies: $missing_deps"
        return 1
      fi
    fi
  fi

  # Display usage information from metadata
  local usage
  usage=$(grep '^usage' "$meta_path" | cut -d'"' -f2)
  echo -e " Usage: $usage"

  # Prompt for arguments
  read -rp "󰛿 Enter arguments for '$name': " args
  echo " Running: $name $args"
  echo

  # Execute the script with the provided arguments
  "$script_path" $args
}
