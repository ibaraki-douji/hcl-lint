package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/mitchellh/colorstring"
)

func main() {
	var parseErr int

	parser := hclparse.NewParser()

	if len(os.Args) == 2 && os.Args[1] == "-" {
		bytes, _ := io.ReadAll(os.Stdin)
		_, diag := parser.ParseHCL(bytes, "stdin")
		if diag.HasErrors() {
			colorstring.Printf("[red]Error parsing stdin: %s\n", diag.Error())
			parseErr = 1
		} else {
			colorstring.Printf("[green]OK!\n")
		}
	} else {
		for i, arg := range os.Args {
			if i == 0 {
				continue
			}
			search := arg
			if info, err := os.Stat(arg); err == nil && info.IsDir() {
				search = fmt.Sprintf("%s/*.tf", arg)
			}
			files, err := filepath.Glob(search)
			if err != nil {
				colorstring.Printf("[red]Error finding files: %s", err)
			}
			for _, filename := range files {
				fmt.Printf("Checking %s ... ", filename)
				file, err := os.ReadFile(filename)
				if err != nil {
					colorstring.Printf("[red]Error reading file: %s\n", err)
					break
				}
				_, diag := parser.ParseHCL(file, filename)
				if diag.HasErrors() {
					colorstring.Printf("[red]Error parsing file: %s\n", diag.Error())
					parseErr = 1
					break
				}
				colorstring.Printf("[green]OK!\n")
			}
		}
	}
	os.Exit(parseErr)
}
