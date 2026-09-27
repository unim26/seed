#  Seed

> The fastest way to snapshot and clone project architectures.

**Seed** is a lightweight CLI tool written in Go that captures the "hollow skeleton" of any existing project and packages it into a portable `.seed` blueprint. You can share this blueprint or use it to instantly scaffold new projects with the exact same directory and file architecture—without carrying over any of the underlying code.

##  Features

*  **Zero Dependencies:** A single, lightning-fast compiled Go binary.
*  **Idempotent by Design:** Seed will *never* overwrite existing files. It safely skips files and folders that already exist, making it completely safe to run in active repositories.
*  **Smart Ignorance:** Automatically ignores heavy dependency folders like `.git`, `node_modules`, etc.
*  **Tamper-Proof Blueprints:** Snapshots are serialized and Base64-encoded into a `.seed` file, protecting the blueprint structure from manual corruption.

## 📖 Usage

Seed is designed to be completely safe to run in active repositories. It uses two simple commands: `snapshot` and `build`.

### 1. Capture a Blueprint
Navigate to an existing project that has a folder architecture you want to reuse, and run:

```bash
seed snapshot
```
This scans your current directory tree, strips out all file contents, ignores dependencies (like .git and node_modules), and saves the empty structural blueprint to a .seed file.

### 2. Scaffold a New Project
Move your newly generated .seed file into a new directory (or share it with your team), and run the build command:

```bash
seed build
```
Seed will instantly recreate all the folders and blank files.


##  Installation

Choose the method that works best for your setup:

**Option 1: Using Go Install (Recommended)**
If you already have Go installed, you can pull and install the latest version directly:
```bash
go install github.com/unim26/seed@latest
```

**Option 2: Compile from Source**
Clone the repository and build the binary manually:
```bash
git clone https://github.com/unim26/seed.git
cd seed
go build -o seed main.go
sudo mv seed /usr/local/bin/
```

**Option 3: Download the Pre-compiled Executable**
1. Head over to the Releases page.
2. Download the executable that matches your operating system (Windows, Mac, or Linux).
3. Place the binary in a folder and add that folder's path to your system's PATH environment variable.


## 📜 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

Build by Abhishek 💓
