# Gopi -- Go Project Initializer

**gopi** is a lightweight Go project initializer built to automate the
setup and boilerplating of new Go projects.\
It reflects **my personal workflow, structure preferences, and project
organization style**, but the generated structure is clean and reusable
--- useful for others who prefer a consistent starting point.

---

## 🚀 Installation

1.  **Clone the repository**

    ```bash
    git clone <repo-url>
    cd gopi
    ```

2.  **Build the binary**

    ```bash
    go build -o bin/gopi .
    ```

    ⚠️ **Important:**\
    The built binary **must be placed inside the `bin/` directory at the
    project root** to ensure correct internal path and resource
    handling.

---

## 🛠️ Usage

You can use the `gopi` binary in two ways:

### 1. Run using its absolute path

```bash
/absolute/path/to/gopi/bin/gopi <args>
```

### 2. Add the binary to your system `PATH`

```bash
export PATH="$PATH:/absolute/path/to/gopi/bin"
```

Then simply call:

```bash
gopi <args>
```

---

## 📦 Initializing a New Go Project

`gopi` can be invoked from **anywhere**. You may:

- Run it **from inside your project root**, or\
- Run it **from any directory** while passing the absolute project
  root path as the first argument.

### Special Case: Using `"."` as `projectPath`

If `projectPath` is `"."`, the **current working directory is treated as
the project root**.

---

## ⚙️ Required Arguments

`gopi` requires **4 mandatory arguments** in the **exact order**:

    gopi <projectPath> <projectName> <author> <goVersion>

---

Argument Description

---

`projectPath` Path to the project root. `"."` means use the current
directory.

`projectName` Name of the Go project

`author` Author name for metadata

`goVersion` Go version for `go mod init` (default: **1.24.10**)

---

### ➤ Using the default Go version

To use the default version, pass an empty string `""` for `goVersion`,
but **the argument must still be included**.

Example:

```bash
gopi . myapp "Moinak Dey" ""
```

⚠️ **Arguments must always be passed in strict order.**

---

## 📁 Example Usage

Initialize a project by specifying an absolute project path:

```bash
gopi /home/moinak/Projects/sampleapp sampleapp "Moinak Dey" 1.24.10
```

Initialize from inside the project root and use the default Go version:

```bash
cd /home/moinak/Projects/sampleapp
gopi . sampleapp "Moinak Dey" ""
```

Or call `gopi` from anywhere if you added the binary to your `PATH`:

```bash
gopi /absolute/path/to/project projectname "Author Name" 1.24.10
```

---

## ✨ About This Tool

`gopi` was built to streamline **my personal Go project setup process**
--- including directory layout, starter code, and config generation ---
based on how I work and maintain project structures.\
While opinionated to my workflow, the output is intentionally minimal
and practical, making it a helpful starting scaffold for other
developers as well.

---

## 📄 License

MIT
