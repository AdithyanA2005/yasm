#!/usr/bin/env bash

function generate_meta_template() {
  cat <<EOF
description = "Add a description here"
tags = ["add", "tags", "here"]
usage = "$1.sh <arguments>"
dependencies = ["add", "dependencies", "here"]
logfile = "logs/$1-\$(date +'%Y%m%d-%H%M%S').log"
EOF
}

function generate_script_template() {
  cat <<EOF
#!/usr/bin/env bash

# $1 script
echo "Hello from $1!"
EOF
}
