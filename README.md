# yasm

_A modern CLI script manager for your toolbox_

> **Note:** yasm is currently under active development. Some features may be incomplete or unstable, and you may encounter bugs. Please use with caution and report any issues you find!

**yasm** is a script manager that allows users to create, edit, run, and manage executable scripts (Bash, Python, etc.) through a consistent CLI. It supports rich script metadata via comments for title, description, tags, usage, and dependencies.

---

## Features

- 📜 Create scripts with shebang and metadata templates
- 🏷️ Add metadata (`@yasm.title`, `@yasm.description`, etc.)
- 🧠 Run scripts with arguments and see usage info
- 🔍 Fuzzy-find scripts with preview
- 🛠️ Edit and delete scripts via CLI
- 📂 Configurable script storage and editor
- 🧩 Support for multiple languages (via shebang)
- 🧪 Extensible with user-defined languages in `config.toml`

---

## Installation

Written in Go. Build from source:

```bash
git clone https://github.com/adithyana2005/yasm
cd yasm
go build -o yasm
```

---

## Usage Examples

```bash
# Create a new script (bash script)
yasm create myscript

# You can specify the language explicitly
yasm create --lang python myscript

# Run a script
yasm run myscript

# List all scripts
yasm list

# Edit a script
yasm edit myscript

# Delete a script interactively
yasm delete

# Show script metadata
yasm info myscript
```

---

## Metadata Annotations

Scripts start with a metadata block:

```bash
#!/usr/bin/env bash

# @yasm.title My Script
# @yasm.description This script does something useful
# @yasm.tags utility network
# @yasm.dependencies curl
```

---

## Configuration

User config file: `~/.config/yasm/config.toml`

You can override:

- `editor` (default editor)
- `script-dir` (where scripts are stored)
- `[languages]` (add or customize language support)

---

## Contributing

Contributions welcome! To run locally:

```bash
go run .
```

---

## License

MIT License. See [LICENSE](LICENSE) for details.
