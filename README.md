# 🛠️ YASM (Yet Another Script Manager)

YASM is a modular, metadata-driven script management system built for reuse, clarity, and extensibility. It helps you organize your scripts with proper metadata, provides fuzzy search UI using `fzf`, checks dependencies before running, and can even be backed up easily with `git`.

---

## 📦 Features

- Organize your scripts with TOML-based metadata
- Fuzzy search and select using `fzf`
- Dependency checking before execution
- Easily git-version your scripts

---

## 🚀 Installation

### 1. Clone the Repo

```bash
git clone https://github.com/AdithyanA2005/yasm
cd yasm
```

### 2. Make the Installer Executable

```bash
chmod +x install.sh
```

### 3. Run the Installer

```bash
./install.sh
```

### 4. Add ~/.local/bin to PATH

Depending on your shell, add this line to the appropriate config file:

- **For bash (~/.bashrc):**

  ```bash
  export PATH="$HOME/.local/bin:$PATH"
  ```

- **For zsh (~/.zshrc):**

  ```bash
  export PATH="$HOME/.local/bin:$PATH"
  ```

- **For fish (~/.config/fish/config.fish):**

  ```bash
  set -Ux PATH $HOME/.local/bin $PATH
  ```

After adding, restart your terminal or source the config file.

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

## 🛠️ Configuration

The config file is located at `~/.config/yasm/config`. You can customize the following

```bash
# Default values
yasm_scripts_dir=$HOME/.dotscripts/scripts
yasm_meta_dir=$HOME/.dotscripts/meta
yasm_editor=$EDITOR # Fallback `nano`
```

- `yasm_scripts_dir`: Directory in which shell files of scripts are stored
- `yasm_meta_dir`: Directory in which metadata files of scripts are stored
- `yasm_editor`: Editor to use for opening scripts and metadata files

---

## 🧠 Metadata (TOML)

Each script has a corresponding `.toml` file in `meta/` with info like:

```toml
# backup.toml
description = "Backup a folder to a destination using rsync"
tags = ["backup", "rsync"]
usage = "backup <source> <destination>"
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
