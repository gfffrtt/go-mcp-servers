# 🌐 MCP Servers Monorepo 🚀

## 📡 What is MCP?

Model Context Protocol (MCP) is a robust, flexible communication framework designed to streamline service interactions and provide a standardized approach to building modular, scalable servers for AI agents.

## 🗂️ Available MCP Servers

### 📂 Filesystem Server

- **Capabilities**:
  - 📖 File reading
  - ✍️ File writing
  - 🔍 File searching
- **Detailed Documentation**: [fs](cmd/fs/README.md)

## 🏗️ Project Structure

```
/mcp-servers
  /cmd            # Main server entry points
    /fs           # Filesystem MCP Server
  /internal       # Internal packages and server implementations
    /fs           # Filesystem server internals
  README.md       # This document
```
