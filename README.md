# 🛠️ DotScripts - Script Manager

DotScripts is a modular, metadata-driven script management system built for reuse, clarity, and extensibility. It helps you organize your scripts with proper metadata, provides fuzzy search UI using `fzf`, checks dependencies before running, and can even be backed up easily with `git`.

---

## 📦 Features

- Organize your scripts with TOML-based metadata
- Fuzzy search and select using `fzf`
- Automatic usage display and argument prompts
- Dependency checking before execution
- Supports Fish shell completions (auto-generated)
- Easily git-version your scripts
- Extendable with your own script templates

---

## 📁 Folder Structure

```
~/.dotscripts/
├── bin/           # Your actual scripts (bash, etc)
├── meta/          # Metadata for each script (TOML)
├── logs/          # Automatically generated logs for each script run
├── lib/           # Core library files (helper functions, fzf UI, etc)
├── completions/   # Auto-generated completions for fish shell
└── dts            # Main executable
```

---

## 🚀 Installation

> In the below example I will use `~/.dotscripts` as the installation directory. You can change it to any directory you prefer. But make sure to update the `PATH` variable accordingly.

### 1. Clone the Repo

```bash
git clone https://github.com/AdithyanA2005/dotscripts ~/.dotscripts
```

### 2. Add to PATH

Add the following line to your `~/.bashrc`, `~/.zshrc`, or `~/.config/fish/config.fish`:

#### Bash / Zsh

```bash
export PATH="$HOME/.dotscripts:$PATH"
```

#### Fish

```fish
set -Ux PATH $HOME/.dotscripts $PATH
```

### 3. Make Executable

```bash
chmod +x ~/.dotscripts/dts
```

Now you can run `dts` from anywhere 🎉

---

## 📄 Usage

```bash
dts             # Launch the fuzzy menu
```

Or directly run a script:

```bash
dts run backup  # Will check for dependencies, prompt for arguments, and run it
```

Create a new script:

```bash
dts new         # Interactive new script creator
```

Rename or delete scripts via the fuzzy menu.

---

## 🧠 Metadata (TOML)

Each script has a corresponding `.toml` file in `meta/` with info like:

```toml
# backup.toml
description = "Backup a folder to a destination using rsync"
tags = ["backup", "rsync"]
usage = "backup.sh <source> <destination>"
dependencies = ["rsync"]
```

---

## ✅ Benefits of Using DotScripts

- Keep your script collection clean and organized
- Quickly find and run scripts using `fzf`
- Auto-check dependencies before running
- Easy to share and reuse across machines
- Plug-and-play with `git` for version control

---

## 💡 Tips for Reusability & Git

1. **Use git** to track your scripts:

```bash
cd ~/.dotscripts
git init
git remote add origin git@github.com:yourusername/dotscripts.git
git add .
git commit -m "Initial commit"
```

2. **Use branches** to manage custom setups (e.g., `work`, `home`, `server`)

3. **Backup often** using `git push`

---

## 🧪 Roadmap (Coming Soon)

- Shell completions (fish ✅, bash/zsh planned)
- Remote script sharing
- Template system for new scripts
- GUI wrapper (maybe 😉)

---

## 🧑‍💻 Author

Made with ❤️ by [Adithyan A](https://github.com/AdithyanA2005)

---

## 📜 License

MIT License. You are free to use, modify, and share.
