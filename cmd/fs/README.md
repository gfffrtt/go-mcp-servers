# 📂 Filesystem MCP Server

A fast and lightweight MCP server written in Go.

## Features

### 📖 Read Tools

- `read_file`: Peek into any file's contents
  - 📍 Retrieve file data with ease
- `read_dir`: Explore directory landscapes
  - 🗂️ List files and subdirectories effortlessly

### 📑 Write Tools

- `write_file`: Craft and save your digital content
  - 📝 Create or overwrite files seamlessly
- `write_dir`: Creates a new directory
  - 🏗️ Safe default to not overwrite directory

### 🔍 Search Capabilities

- `search`: Search for term in directory recursively
  - 🕵️ Powerful content discovery using [`ripgrep`](https://github.com/BurntSushi/ripgrep)

## 🚀 Quick Start

### Prerequisites

- 🐹 Go 1.23+
- 🔍 [Ripgrep](https://github.com/BurntSushi/ripgrep)

### Installation

```bash
git clone go-mcp-servers

make fs
```

### Usage

> 📌 Claude desktop example

```json
{
  "mcpServers": {
    "fs": {
      "command": "/path/to/repo/go-mcp-servers/build/fs",
      "args": ["/path/to/folder"]
    }
  }
}
```
