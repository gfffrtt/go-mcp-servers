package git

import (
	"context"
	"os/exec"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

type Diff struct {
	Server *server.MCPServer
	Path   string
}

func NewDiff(server *server.MCPServer, path string) *Diff {
	return &Diff{
		Server: server,
		Path:   path,
	}
}

func (d *Diff) Diff() {
	diffTool := mcp.NewTool(
		"diff",
		mcp.WithDescription("Get the diff of the current git repository"),
	)
	d.Server.AddTool(diffTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		cmd := exec.Command("git", "diff")
		cmd.Dir = d.Path
		out, err := cmd.Output()
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText(string(out)), nil
	})
}

func (d *Diff) AddTools() {
	d.Diff()
}
