package fs

import (
	"context"
	"fmt"
	"os"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

type Read struct {
	Server *server.MCPServer
	Path   string
}

func NewRead(server *server.MCPServer, path string) *Read {
	return &Read{
		Server: server,
		Path:   path,
	}
}

func (r *Read) File() {
	readFileTool := mcp.NewTool(
		"read_file",
		mcp.WithDescription("Reads a file from the filesystem with the given path"),
		mcp.WithString("path", mcp.Required(), mcp.Description(fmt.Sprintf("The path to the file to read. The path should be the relative path from the current working directory %s", r.Path))),
	)
	r.Server.AddTool(readFileTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		path := request.Params.Arguments["path"].(string)

		content, err := os.ReadFile(path)
		if err != nil {
			return nil, err
		}

		return mcp.NewToolResultText(string(content)), nil
	})
}

func (r *Read) Dir() {
	readDirTool := mcp.NewTool(
		"read_dir",
		mcp.WithDescription("Reads a directory from the filesystem with the given path"),
		mcp.WithString("path", mcp.Required(), mcp.Description(fmt.Sprintf("The path to the directory to read. The path should be the relative path from the current working directory %s", r.Path))),
	)
	r.Server.AddTool(readDirTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		path := request.Params.Arguments["path"].(string)

		files, err := os.ReadDir(path)
		if err != nil {
			return nil, err
		}

		var content string
		for _, file := range files {
			content += file.Name() + "\n"
		}

		return mcp.NewToolResultText(content), nil
	})
}

func (r *Read) AddTools() {
	r.File()
	r.Dir()
}
