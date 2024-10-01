package main

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
)

const (
	practiceDir  = "practice"
	templatesDir = "templates"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: kata-manager {new-day|copy-kata <template_path> [target_day]|last-day}")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "new-day":
		if err := generateNewDay(); err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
	case "copy-kata":
		if len(os.Args) < 3 {
			fmt.Println("Usage: kata-manager copy-kata <template_path> [target_day]")
			os.Exit(1)
		}
		templatePath := os.Args[2]
		targetDay := ""
		if len(os.Args) > 3 {
			targetDay = os.Args[3]
		}
		if err := copyKata(templatePath, targetDay); err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
	case "last-day":
		lastDay, err := findLastDay()
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		fmt.Println(lastDay)
	default:
		fmt.Println("Unknown command. Use 'new-day', 'copy-kata', or 'last-day'.")
		os.Exit(1)
	}
}

func getLatestDay(entries []fs.DirEntry) int {

	latestDay := 0
	for _, entry := range entries {
		if entry.IsDir() && len(entry.Name()) > 4 && entry.Name()[:4] == "day_" {
			if dayNum, err := strconv.Atoi(entry.Name()[4:]); err == nil && dayNum > latestDay {
				latestDay = dayNum
			}
		}
	}
	return latestDay
}

func generateNewDay() error {
	entries, err := os.ReadDir(practiceDir)

	if err != nil {
		if os.IsNotExist(err) {
			if err := os.Mkdir(practiceDir, 0755); err != nil {
				return fmt.Errorf("failed to create practice directory: %w", err)
			}
		} else {
			return fmt.Errorf("failed to read practice directory: %w", err)
		}
	}

	latestDay := getLatestDay(entries)

	newDay := fmt.Sprintf("day_%d", latestDay+1)
	newDayPath := filepath.Join(practiceDir, newDay)
	if err := os.Mkdir(newDayPath, 0755); err != nil {
		return fmt.Errorf("failed to create new day directory: %w", err)
	}

	fmt.Printf("Created new practice day: %s\n", newDay)
	return nil
}

func copyKata(templatePath, targetDay string) error {
	// Validate template path
	templateInfo, err := os.Stat(templatePath)
	if err != nil {
		return fmt.Errorf("template path does not exist: %w", err)
	}
	if !templateInfo.IsDir() {
		return fmt.Errorf("template path is not a directory")
	}

	// Determine target day
	if targetDay == "" {
		entries, err := os.ReadDir(practiceDir)
		if err != nil {
			return fmt.Errorf("failed to read practice directory: %w", err)
		}

		latestDay := getLatestDay(entries)

		if latestDay == 0 {
			return fmt.Errorf("no practice days found, please generate a new day first")
		}
		targetDay = fmt.Sprintf("day_%d", latestDay)
	}

	targetDayPath := filepath.Join(practiceDir, targetDay)
	if _, err := os.Stat(targetDayPath); os.IsNotExist(err) {
		return fmt.Errorf("practice day does not exist: %s", targetDay)
	}

	kataName := filepath.Base(templatePath)
	targetPath := filepath.Join(targetDayPath, kataName)

	if _, err := os.Stat(targetPath); !os.IsNotExist(err) {
		return fmt.Errorf("kata '%s' already exists in %s", kataName, targetDay)
	}

	if err := copyDir(templatePath, targetPath); err != nil {
		return fmt.Errorf("failed to copy kata: %w", err)
	}

	fmt.Printf("Copied kata '%s' to %s\n", kataName, targetDay)
	return nil
}

func copyDir(src, dst string) error {
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip files starting with underscore or README.md files
		baseName := filepath.Base(path)
		if baseName[0] == '_' || baseName == "README.md" {
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		dstPath := filepath.Join(dst, relPath)

		if info.IsDir() {
			return os.MkdirAll(dstPath, info.Mode())
		}

		return copyFile(path, dstPath)
	})
}

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Close()
}

func findLastDay() (string, error) {
	entries, err := os.ReadDir(practiceDir)
	if err != nil {
		return "", fmt.Errorf("failed to read practice directory: %w", err)
	}

	var lastDay string
	var lastNum int
	for _, entry := range entries {
		if entry.IsDir() && len(entry.Name()) > 4 && entry.Name()[:4] == "day_" {
			if num, err := strconv.Atoi(entry.Name()[4:]); err == nil {
				if num > lastNum {
					lastNum = num
					lastDay = entry.Name()
				}
			}
		}
	}

	if lastDay == "" {
		return "", fmt.Errorf("no practice days found")
	}

	return filepath.Join(practiceDir, lastDay), nil
}
