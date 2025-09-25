package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/hashicorp/hcl/v2/hclparse"
)

type Config struct {
	Recursive bool
	FileTypes []string
}

func main() {
	parser := hclparse.NewParser()

	// Define flags
	var (
		recursive     = flag.Bool("r", false, "recursively search subdirectories")
		recursiveLong = flag.Bool("recursive", false, "recursively search subdirectories")
		fileType      = flag.String("t", "tf", "file extensions to process (comma-separated, e.g., tf,hcl,tfvars)")
		fileTypeLong  = flag.String("type", "tf", "file extensions to process (comma-separated, e.g., tf,hcl,tfvars)")
		help          = flag.Bool("h", false, "show help")
		helpLong      = flag.Bool("help", false, "show help")
		version       = flag.Bool("v", false, "show version")
		versionLong   = flag.Bool("version", false, "show version")
	)

	// Custom usage function
	flag.Usage = PrintUsage

	// Parse flags
	flag.Parse()

	// Handle help and version
	if *help || *helpLong {
		PrintUsage()
		os.Exit(0)
	}

	if *version || *versionLong {
		PrintVersion()
		os.Exit(0)
	}

	// Get remaining arguments after flags
	args := flag.Args()

	// Handle stdin input
	if len(args) == 1 && args[0] == "-" {
		exitCode := ProcessStdin(parser)
		os.Exit(exitCode)
	}

	// Check if no arguments provided
	if len(args) == 0 {
		PrintShortUsage()
		os.Exit(1)
	}

	// Create config
	config := Config{
		Recursive: *recursive || *recursiveLong,
	}

	// Parse file types
	fileTypeStr := *fileType
	if *fileTypeLong != "tf" {
		fileTypeStr = *fileTypeLong
	}
	config.FileTypes = parseFileTypes(fileTypeStr)

	// Expand globs
	expandedArgs, err := expandGlobs(args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error expanding globs: %v\n", err)
		os.Exit(1)
	}

	if len(expandedArgs) == 0 {
		PrintShortUsage()
		os.Exit(1)
	}

	exitCode := ProcessFiles(parser, expandedArgs, config)
	os.Exit(exitCode)
}

func parseFileTypes(typeStr string) []string {
	types := strings.Split(typeStr, ",")
	for i, t := range types {
		t = strings.TrimSpace(t)
		if !strings.HasPrefix(t, ".") {
			t = "." + t
		}
		types[i] = t
	}
	return types
}
