package main

import (
	"fmt"
	"os"
)

func PrintUsage() {
	fmt.Printf(`HCL Lint - A command-line tool for validating HCL (HashiCorp Configuration Language) files

USAGE:
    %s [OPTIONS] [FILES...]
    %s [OPTIONS] [DIRECTORIES...]
    cat file.hcl | %s -

DESCRIPTION:
    HCL Lint validates the syntax of HCL files including Terraform configurations.
    It can process individual files, directories (with optional recursive scanning),
    or read from stdin.

OPTIONS:
    -h, --help              Show this help message and exit
    -v, --version           Show version information and exit
    -r, --recursive         Recursively search subdirectories for files
    -t, --type <extensions> File extensions to process (comma-separated, default: tf)
                           Examples: tf, hcl, tfvars, tf,hcl, hcl,tfvars
    -                       Read HCL content from stdin

ARGUMENTS:
    FILES...        One or more HCL or Terraform files to validate
    DIRECTORIES...  One or more directories to scan for files

EXAMPLES:
    # Validate a single file
    %s main.tf

    # Validate all .tf files in a directory (non-recursive)
    %s ./terraform/

    # Validate all .tf files in a directory and subdirectories
    %s -r ./terraform/

    # Validate .hcl files instead of .tf files
    %s -t hcl ./config/

    # Validate multiple file types recursively
    %s -r -t tf,hcl,tfvars ./project/

    # Validate multiple files and directories
    %s main.tf ./modules variables.tf

    # Read from stdin
    cat config.tf | %s -

EXIT CODES:
    0    All files are syntactically valid
    1    One or more files contain syntax errors

SUPPORTED FILE TYPES:
    - .tf (Terraform configuration files)
    - .hcl (HashiCorp Configuration Language files)
    - .tfvars (Terraform variable files)
    - Any file when explicitly specified (content will be parsed as HCL)

`, os.Args[0], os.Args[0], os.Args[0], os.Args[0], os.Args[0], os.Args[0], os.Args[0], os.Args[0], os.Args[0], os.Args[0])
}

func PrintVersion() {
	fmt.Printf("HCL Lint v1.0.0\n")
	fmt.Printf("Built with Go and HashiCorp's HCL parser\n")
	fmt.Printf("Repository: https://github.com/ibaraki-douji/hcl-lint\n")
}

func PrintShortUsage() {
	fmt.Printf("Usage: %s [OPTIONS] [FILES/DIRECTORIES...]\n", os.Args[0])
	fmt.Printf("Use '%s --help' for more information.\n", os.Args[0])
}
