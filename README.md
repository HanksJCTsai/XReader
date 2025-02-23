```
__   __ _____                    _
\ \ / /|  __ \                  | |
 \ V / | |__) |  ___   __ _   __| |  ___  _ __
  > <  |  _  /  / _ \ / _` | / _` | / _ \| '__|
 / . \ | | \ \ |  __/| (_| || (_| ||  __/| |
/_/ \_\|_|  \_\ \___| \__,_| \__,_| \___||_|

```
<div align="center">
  <a href="https://golang.org/"><img src="https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go" alt="Go Version"></a>
  <a href="https://github.com/HanksJCTsai/XReader/blob/main/LICENSE"><img src="https://img.shields.io/github/license/HanksJCTsai/XReader" alt="License"></a>
  <a href="https://github.com/HanksJCTsai/XReader/releases"><img src="https://img.shields.io/github/v/release/HanksJCTsai/XReader" alt="GitHub Release"></a>
</div>

# 🚀 XReader

## 🌟 Introduction

🎉 **XReader** is a sleek command-line tool crafted in Go for *efficient file reading*. Perfect for quick file operations or sharpening your Go skills. Let’s make file handling a breeze!

---

## ✨ Features

Here’s what XReader brings to the table:
- **🚀 Read and display file contents.**
- ⚙️ Read files line-by-line and display line numbers.
- 🛠 Display file metadata (size, last modified time).
- 🛠️ Support multi-file processing and content filtering.
- 🛠 Text convert to json.
---

## 📦 Installation

#### Prerequisites
- [Go](https://go.dev/dl/) (1.21 or higher recommended).
- [Git](https://git-scm.com/downloads) installed.
#### Steps
1. **Clone the repository:**
   ```bash
   cd your_want_store_source_code_dir or mkdir your_want_store_source_code_dir
   git clone https://github.com/HanksJCTsai/XReader.git
   cd XReader
   ```
2. **Build the tool:**
   ```bash
   go build -o XReader ./cmd
   ```
3. ***(Optional) Add to PATH:***
   ```bash
   sudo mv XReader /usr/local/bin/
   ```
---

## 🚀 Usage

#### Flags
Below are the available command-line flags:

| Flag           | Description                 | Default       |
|---------------|-----------------------------|--------------|
| `-file`       | 🎯 Path to the file to read | None         |

#### Example
   For a file `test.txt` with the content: 📖
   ```text
   Hello, Go!
   ```
   Run the command: 🚀
   ```bash
   XReader -file test.txt
   ```
   Output: 🎉
   ```text
   File content:
   Hello, Go!
   ```
---

## 🧪 Testing

### Run tests:
```bash
go test ./internal/fileops
```

---

## 📜 License

**[MIT License](https://opensource.org/licenses/MIT)**  
A permissive open-source license allowing flexibility and freedom.
---
## 📩 Get in Touch
Got questions or ideas?
- **[Open an Issue](https://github.com/HanksJCTsai/XReader/issues)** – Let’s discuss it!
- **Email**: [Hanks' Mail](mailto:u2_u2@msn.com)
---
<div align="center"> <em>Last updated: February 22, 2025</em><br> <em>Built with ❤️ by Hanks Tsai</em> </div>