package fs

import (
	"context"
	"fmt"
	"os/exec"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

type Search struct {
	Server *server.MCPServer
	Path   string
}

func NewSearch(s *server.MCPServer, path string) *Search {
	return &Search{
		Server: s,
		Path:   path,
	}
}

func (s *Search) Search() {
	searchTool := mcp.NewTool("search",
		mcp.WithDescription("Search for files in the directoy of the path provided"),
		mcp.WithString("path", mcp.Required(), mcp.Description(fmt.Sprintf("The path to the directory to search. The path should be the relative path from the current working directory %s", s.Path))),
		mcp.WithString("query", mcp.Required(), mcp.Description("The query to search for")),
	)
	s.Server.AddTool(searchTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		query := request.Params.Arguments["query"].(string)
		path := request.Params.Arguments["path"].(string)

		cmd := exec.Command("rg", "-l", query, path)

		out, err := cmd.CombinedOutput()
		if err != nil {
			return nil, err
		}

		return mcp.NewToolResultText(string(out)), nil
	})
}

func (s *Search) AddTools() {
	s.Search()
}
