#!/usr/bin/env bash

set -e

INSTALL_DIR="$HOME/.local/bin"
INSTALL_PATH="$INSTALL_DIR/yasm"

# Get script's directory
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
YASM_SOURCE="$SCRIPT_DIR/yasm"

mkdir -p "$INSTALL_DIR"
cp "$YASM_SOURCE" "$INSTALL_PATH"
chmod +x "$INSTALL_PATH"

echo "✅ Installed yasm to $INSTALL_PATH"
echo "👉 Make sure $INSTALL_DIR is in your PATH."
