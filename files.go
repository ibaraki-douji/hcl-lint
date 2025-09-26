package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/mitchellh/colorstring"
)

func ProcessFiles(parser *hclparse.Parser, args []string, config Config) int {
	var parseErr int

	for _, arg := range args {
		if err := processPath(parser, arg, config); err != nil {
			parseErr = 1
		}
	}

	return parseErr
}

func processPath(parser *hclparse.Parser, path string, config Config) error {
	info, err := os.Stat(path)
	if err != nil {
		colorstring.Printf("[red]Error accessing path '%s': %s\n", path, err)
		return err
	}

	if info.IsDir() {
		return processDirectory(parser, path, config)
	}

	return processFile(parser, path, config)
}

func processDirectory(parser *hclparse.Parser, dirPath string, config Config) error {
	var files []string

	if config.Recursive {
		// Walk directory recursively
		err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			// Check if this path is excluded
			if isExcluded(path, config.Exclude) {
				if info.IsDir() {
					return filepath.SkipDir
				}
				return nil
			}

			if !info.IsDir() && hasValidExtension(path, config.FileTypes) {
				files = append(files, path)
			}
			return nil
		})
		if err != nil {
			colorstring.Printf("[red]Error walking directory '%s': %s\n", dirPath, err)
			return err
		}
	} else {
		// Only search current directory
		for _, ext := range config.FileTypes {
			searchPattern := filepath.Join(dirPath, "*"+ext)
			matches, err := filepath.Glob(searchPattern)
			if err != nil {
				colorstring.Printf("[red]Error finding %s files in directory '%s': %s\n", ext, dirPath, err)
				return err
			}
			files = append(files, matches...)
		}
	}

	if len(files) == 0 {
		extensionsStr := strings.Join(config.FileTypes, ", ")
		recursiveStr := ""
		if config.Recursive {
			recursiveStr = " (recursively)"
		}
		colorstring.Printf("[yellow]No %s files found in directory '%s'%s\n", extensionsStr, dirPath, recursiveStr)
		return nil
	}

	var hasError bool
	for _, filename := range files {
		if err := processFile(parser, filename, config); err != nil {
			hasError = true
		}
	}

	if hasError {
		return fmt.Errorf("errors found in directory %s", dirPath)
	}

	return nil
}

func processFile(parser *hclparse.Parser, filename string, config Config) error {
	// Check if file is excluded
	if isExcluded(filename, config.Exclude) {
		return nil
	}

	fmt.Printf("Checking %s ... ", filename)

	file, err := os.ReadFile(filename)
	if err != nil {
		colorstring.Printf("[red]Error reading file: %s\n", err)
		return err
	}

	_, diag := parser.ParseHCL(file, filename)
	if diag.HasErrors() {
		colorstring.Printf("[red]Error parsing file: %s\n", diag.Error())
		return fmt.Errorf("parse error in %s", filename)
	}

	colorstring.Printf("[green]OK!\n")
	return nil
}

func hasValidExtension(filename string, validExtensions []string) bool {
	ext := filepath.Ext(filename)
	for _, validExt := range validExtensions {
		if ext == validExt {
			return true
		}
	}
	return false
}

func expandGlobs(args []string) ([]string, error) {
	var expanded []string

	for _, arg := range args {
		if strings.ContainsAny(arg, "*?[]") {
			matches, err := filepath.Glob(arg)
			if err != nil {
				return nil, fmt.Errorf("error expanding glob pattern '%s': %v", arg, err)
			}
			if len(matches) == 0 {
				colorstring.Printf("[yellow]No files match pattern '%s'\n", arg)
				continue
			}
			expanded = append(expanded, matches...)
		} else {
			expanded = append(expanded, arg)
		}
	}

	return expanded, nil
}

// isExcluded checks if the given path matches any of the exclude patterns
func isExcluded(path string, excludePatterns []string) bool {
	if len(excludePatterns) == 0 {
		return false
	}

	// Clean the path to handle both relative and absolute paths consistently
	cleanPath := filepath.Clean(path)

	// Also check without the leading "./"
	relPath := strings.TrimPrefix(cleanPath, "./")

	for _, pattern := range excludePatterns {
		pattern = strings.TrimSpace(pattern)
		if pattern == "" {
			continue
		}

		// Clean the pattern
		cleanPattern := filepath.Clean(pattern)

		// Check if the path exactly matches the pattern
		if cleanPath == cleanPattern || relPath == cleanPattern {
			return true
		}

		// Check if the path is under an excluded directory
		if strings.HasPrefix(cleanPath+"/", cleanPattern+"/") ||
			strings.HasPrefix(relPath+"/", cleanPattern+"/") {
			return true
		}

		// For absolute paths, also check if the path starts with the pattern
		if filepath.IsAbs(cleanPath) && strings.HasPrefix(cleanPath, "/"+cleanPattern) {
			return true
		}
	}

	return false
}
