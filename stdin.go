package main

import (
	"io"
	"os"

	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/mitchellh/colorstring"
)

func ProcessStdin(parser *hclparse.Parser) int {
	bytes, err := io.ReadAll(os.Stdin)
	if err != nil {
		colorstring.Printf("[red]Error reading from stdin: %s\n", err)
		return 1
	}

	_, diag := parser.ParseHCL(bytes, "stdin")
	if diag.HasErrors() {
		colorstring.Printf("[red]Error parsing stdin: %s\n", diag.Error())
		return 1
	}

	colorstring.Printf("[green]OK!\n")
	return 0
}
