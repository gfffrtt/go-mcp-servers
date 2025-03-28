package main

import (
	"fmt"
	"go-mcp-servers/internal/fs"
	"os"

	"github.com/mark3labs/mcp-go/server"
)

func main() {
	s := server.NewMCPServer(
		"filesystem",
		"1.0.0",
		server.WithResourceCapabilities(true, true),
		server.WithLogging(),
	)

	path := os.Args[1]

	fs.NewRead(s, path).AddTools()
	fs.NewWrite(s, path).AddTools()
	fs.NewSearch(s, path).AddTools()

	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
