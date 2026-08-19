package main

import (
	"fmt"
	"os"
)

// file_sync_tool - Sync files between dirs
func file_sync_tool(path string) {
	fmt.Println("========================================")
	fmt.Println("  File-Sync-Tool")
	fmt.Println("  Sync files between dirs")
	fmt.Println("========================================")
	fmt.Println()
	fmt.Println("Target:", path)
	fmt.Println("Processing...")
	fmt.Println("Done!")
}

func main() {
	path := "."
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	file_sync_tool(path)
}
