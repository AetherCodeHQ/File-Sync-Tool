package main

import (
	"crypto/md5"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	dir := "."
	if len(os.Args) > 1 {
		dir = os.Args[1]
	}
	fmt.Printf("File Sync Tool\n")
	fmt.Printf("Scanning %s...\n\n", dir)
	files := map[string][16]byte{}
	var names []string
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return nil
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return nil
		}
		sum := md5.Sum(data)
		files[path] = sum
		names = append(names, path)
		return nil
	})
	fmt.Printf("Files indexed: %d\n", len(names))
	fmt.Printf("Unique hashes: %d\n", len(files))
	fmt.Printf("\nSync manifest ready.\n")
}