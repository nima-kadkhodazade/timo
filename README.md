# Timo

<p align="center">
  <strong>A simple, fast and lightweight task manager for your terminal.</strong>
</p>

<p align="center">
  Manage your tasks without leaving the command line.
</p>

<p align="center">
  <a href="https://github.com/nima-kadkhodazade/timo/releases">
    <img src="https://img.shields.io/github/v/release/nima-kadkhodazade/timo?style=flat-square" alt="Latest Release">
  </a>
  <a href="https://github.com/nima-kadkhodazade/timo">
    <img src="https://img.shields.io/github/stars/nima-kadkhodazade/timo?style=flat-square" alt="GitHub Stars">
  </a>
  <a href="https://github.com/nima-kadkhodazade/timo/blob/main/LICENSE">
    <img src="https://img.shields.io/github/license/nima-kadkhodazade/timo?style=flat-square" alt="License">
  </a>
  <img src="https://img.shields.io/badge/built%20with-Go-00ADD8?style=flat-square&logo=go" alt="Built with Go">
</p>

---

## ✨ Why Timo?

Timo is a minimal command-line task manager built with Go.

It is designed for people who want to keep track of their tasks without opening a full productivity application.

No account.

No server.

No unnecessary UI.

Just your terminal and your tasks.

```bash
timo add "Learn Go"
timo list
timo mark-in-progress 1
timo mark-done 1
```

---

## 🚀 Features

- Create tasks from the command line
- List all tasks
- Update task descriptions
- Delete tasks
- Change task status
- Relative creation time (`5 minutes ago`, `2 hours ago`, etc.)
- Persistent local storage
- Lightweight JSON-based storage
- Native executable with no runtime dependencies
- Linux and macOS binaries
- Automated builds through GitHub Actions

### Task statuses

Timo currently supports three task statuses:

- `todo`
- `in-progress`
- `done`

---

## 📦 Installation

### Pre-built binaries

Download the latest release from:

**[GitHub Releases](https://github.com/nima-kadkhodazade/timo/releases)**

Available builds:

| Platform | Architecture |
|----------|--------------|
| Linux | AMD64 |
| Linux | ARM64 |
| macOS | AMD64 |
| macOS | ARM64 |

> Installation scripts and package-manager support are planned for a future release.

---

## ⚡ Quick Start

After installing the Timo binary, make sure it is available in your `PATH`.

Then:

```bash
timo add "Learn Go"
```

List your tasks:

```bash
timo list
```

Example:

```text
ID: 1 | Description: Learn Go | Status: todo | Created: 2 minutes ago
```

---

## 📝 Commands

### Add a task

```bash
timo add "Learn Go"
```

Creates a new task with:

- A unique ID
- `todo` status
- Creation time
- Update time

### List tasks

```bash
timo list
```

Displays all saved tasks.

Example:

```text
ID: 1 | Description: Learn Go | Status: todo | Created: 5 minutes ago
ID: 2 | Description: Build a CLI | Status: in-progress | Created: 1 hour ago
```

### Update a task

```bash
timo update 1 "Learn Go and build a CLI"
```

Updates the description of task `1`.

### Delete a task

```bash
timo delete 1
```

Deletes the task with ID `1`.

### Change task status

Mark a task as **todo**:

```bash
timo mark-todo 1
```

Mark a task as **in progress**:

```bash
timo mark-in-progress 1
```

Mark a task as **done**:

```bash
timo mark-done 1
```

---

## 💾 Data Storage

Timo stores your tasks locally as JSON.

The data directory is created automatically when Timo is first used.

On Linux and macOS, the current storage location is:

```text
~/.local/share/timo/tasks.json
```

Timo does not require a database, server, or internet connection to manage your tasks.

Your task data stays on your machine.

---

## 🛠️ Building From Source

Timo is written in Go.

### Requirements

- Go 1.24.4 or newer

Clone the repository:

```bash
git clone https://github.com/nima-kadkhodazade/timo.git
cd timo
```

Build:

```bash
go build -o timo
```

Run:

```bash
./timo list
```

Or install it into your Go binary directory:

```bash
go install .
```

---

## 🧪 Development

Run the project:

```bash
go run . list
```

Run tests:

```bash
go test ./...
```

Build for Linux:

```bash
GOOS=linux GOARCH=amd64 go build -o timo-linux-amd64
```

Build for Linux ARM64:

```bash
GOOS=linux GOARCH=arm64 go build -o timo-linux-arm64
```

Build for macOS Intel:

```bash
GOOS=darwin GOARCH=amd64 go build -o timo-darwin-amd64
```

Build for macOS Apple Silicon:

```bash
GOOS=darwin GOARCH=arm64 go build -o timo-darwin-arm64
```

---

## 🤖 CI/CD

Timo uses GitHub Actions to automate release builds.

When a version tag is pushed:

```bash
git tag v0.1.1
git push origin v0.1.1
```

GitHub Actions automatically:

1. Checks out the source code
2. Sets up Go
3. Runs the test suite
4. Builds binaries for supported platforms
5. Creates a GitHub Release
6. Uploads the generated binaries

Supported release targets:

```text
linux/amd64
linux/arm64
darwin/amd64
darwin/arm64
```

---

## 🗺️ Roadmap

Timo is actively evolving.

### Current

- [x] Add tasks
- [x] List tasks
- [x] Update tasks
- [x] Delete tasks
- [x] Task statuses
- [x] Relative task creation time
- [x] Local JSON storage
- [x] Cross-platform builds
- [x] GitHub Releases
- [x] Automated release workflow

### Next

- [ ] Cleaner terminal output
- [ ] Input validation improvements
- [ ] Automated tests
- [ ] Better error handling
- [ ] Release archives
- [ ] Checksums for releases
- [ ] Linux installer
- [ ] macOS installer
- [ ] Easier installation through package managers

---

## 🎯 Philosophy

Timo is intentionally simple.

The goal is not to become another complicated project-management platform.

The idea is:

> **Open your terminal. Add a task. Get things done.**

---

## 🤝 Contributing

Contributions, ideas and feedback are welcome.

If you find a bug or have an idea for a feature, feel free to open an issue.

You can also fork the repository and submit a pull request.

---

## 📄 License

This project is open source.

See the [LICENSE](LICENSE) file for details.

---

## ⭐ Support

If Timo is useful to you, consider giving the project a ⭐ on GitHub.

It helps the project grow and motivates future development.

---

<p align="center">
  Built with ❤️ and Go.
</p>
