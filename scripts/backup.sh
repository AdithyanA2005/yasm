#!/usr/bin/env bash

# name: backup.sh
# desc: Backup a folder to a destination using rsync
# usage: backup.sh <source> <destination>
# tags: backup, rsync
# deps: rsync

# Exit immediately if any command fails to ensure safe and predictable behavior
set -e

SRC="$1"
DEST="$2"

if [[ -z "$SRC" || -z "$DEST" ]]; then
  echo "Usage: $0 <source> <destination>"
  exit 1
fi

rsync -a --info=progress2 "$SRC" "$DEST" | tee "$LOG"
echo "Backup completed. Log saved to $LOG"
