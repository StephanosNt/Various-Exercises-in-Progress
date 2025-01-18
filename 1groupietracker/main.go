package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// Define the root directory
	rootDir := "."

	// Walk through the directory tree
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// If it's a directory, create an empty.txt file
		if info.IsDir() {
			emptyFilePath := filepath.Join(path, "empty.txt")
			if _, err := os.Stat(emptyFilePath); os.IsNotExist(err) {
				// Create the empty file
				file, err := os.Create(emptyFilePath)
				if err != nil {
					return fmt.Errorf("failed to create file %s: %w", emptyFilePath, err)
				}
				file.Close()
				fmt.Println("Created:", emptyFilePath)
			} else {
				fmt.Println("File already exists:", emptyFilePath)
			}
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
