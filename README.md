# 🛠️ YASM (Yet Another Script Manager)

YASM is a modular, metadata-driven script management system built for reuse, clarity, and extensibility. It helps you organize your scripts with proper metadata, provides fuzzy search UI using `fzf`, checks dependencies before running, and can even be backed up easily with `git`.

---

## 📦 Features

- Organize your scripts with TOML-based metadata
- Fuzzy search and select using `fzf`
- Dependency checking before execution
- Easily git-version your scripts

---

## 📁 Folder Structure

```
~/.dotscripts/
├── bin/           # Your actual scripts (bash, etc)
├── meta/          # Metadata for each script (TOML)
├── lib/           # Core library files (helper functions, fzf UI, etc)
└── yasm           # Main executable
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
chmod +x ~/.dotscripts/yasm
```

Now you can run `yasm` from anywhere 🎉

---

## 📄 Usage

```bash
yasm             # Launch the fuzzy menu
```

Or directly run a script:

```bash
yasm run backup  # Will check for dependencies, prompt for arguments, and run it
```

Create a new script:

```bash
yasm new         # Interactive new script creator
```

Show help menu

```bash
yasm help
```

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

## ✅ Benefits of Using YASM

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
