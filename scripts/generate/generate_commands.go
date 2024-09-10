package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	root := "templates"
	output := "katas.txt"

	commands, err := generateCommands(root)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	err = os.WriteFile(output, []byte(strings.Join(commands, "\n")), 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		os.Exit(1)
	}

	fmt.Println("Commands written to", output)
}

func generateCommands(root string) ([]string, error) {
	var commands []string

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() && path != root {
			relPath, err := filepath.Rel(root, path)
			if err != nil {
				return err
			}

			// Skip the top-level directories
			if !strings.Contains(relPath, string(os.PathSeparator)) {
				return nil
			}

			command := fmt.Sprintf("practice_kata copy-kata %s", path)
			commands = append(commands, command)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return commands, nil
}