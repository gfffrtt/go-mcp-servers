package fs

import (
	"context"
	"os"
	"path/filepath"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

type Write struct {
	Server *server.MCPServer
	Path   string
}

func NewWrite(server *server.MCPServer, path string) *Write {
	return &Write{
		Server: server,
		Path:   path,
	}
}

func (w *Write) File() {
	writeFileTool := mcp.NewTool(
		"write_file",
		mcp.WithDescription("Writes a file to the filesystem with the given path, it will overwrite the file if it exists"),
		mcp.WithString("path", mcp.Required(), mcp.Description("The path to the file to write")),
		mcp.WithString("content", mcp.Required(), mcp.Description("The content to write to the file")),
	)
	w.Server.AddTool(writeFileTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		path := request.Params.Arguments["path"].(string)
		content := request.Params.Arguments["content"].(string)
		err := os.WriteFile(filepath.Join(w.Path, path), []byte(content), 0644)
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText("Success"), nil
	})
}

func (w *Write) Dir() {
	writeDirTool := mcp.NewTool(
		"write_dir",
		mcp.WithDescription("Writes a directory to the filesystem with the given path, it will not overwrite the directory if it exists"),
		mcp.WithString("path", mcp.Required(), mcp.Description("The path to the directory to write")),
	)
	w.Server.AddTool(writeDirTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		path := request.Params.Arguments["path"].(string)
		err := os.Mkdir(filepath.Join(w.Path, path), 0755)
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText("Success"), nil
	})
}

func (w *Write) AddTools() {
	w.File()
	w.Dir()
}
