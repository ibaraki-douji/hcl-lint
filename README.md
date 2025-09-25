# HCL Lint

A simple command-line tool for validating HCL (HashiCorp Configuration Language) files, including Terraform configurations.

Originally cloned from [github.com/n2ux/hcllint](https://github.com/n2ux/hcllint)

## Features

- Validates HCL syntax in `.tf` and `.hcl` files
- Supports reading from stdin or file/directory paths
- Colored output for better readability
- Recursive directory scanning for `.tf` files
- Built with Go using HashiCorp's official HCL parser

## Installation

### Using go install
```sh
go install github.com/ibaraki-douji/hcl-lint@latest
```

### Building from source
```sh
git clone https://github.com/ibaraki-douji/hcl-lint.git
cd hcl-lint
go build -o hcl-lint .
```

## Usage

### Validate files from stdin
```sh
cat config.tf | hcl-lint -
```

### Validate specific files
```sh
hcl-lint main.tf variables.tf
```

### Validate all .tf files in a directory
```sh
hcl-lint /path/to/terraform/configs
```

### Validate multiple files and directories
```sh
hcl-lint main.tf ./modules ./environments
```

## Examples

**Valid HCL file:**
```
$ hcl-lint tests/valid.hcl
Checking tests/valid.hcl ... OK!
```

**Invalid HCL file:**
```
$ hcl-lint tests/invalid.hcl
Checking tests/invalid.hcl ... Error parsing file: tests/invalid.hcl:13,1-2: Argument or block definition required; An argument or block definition is required here. To set an argument, use the equals sign "=" to introduce the argument value.
```

## Exit Codes

- `0`: All files are valid
- `1`: One or more files contain syntax errors

## Requirements

- Go 1.24 or later
- Dependencies are managed via Go modules