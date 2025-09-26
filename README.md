# HCL Lint

A simple command-line tool for validating HCL (HashiCorp Configuration Language) files, including Terraform configurations.

Originally cloned from [github.com/n2ux/hcllint](https://github.com/n2ux/hcllint)

## Features

- Validates HCL syntax in `.tf`, `.hcl`, `.tfvars`, and other HCL files
- Supports reading from stdin or file/directory paths
- Colored output for better readability
- Optional recursive directory scanning
- Configurable file type filtering
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

```
hcl-lint [OPTIONS] [FILES/DIRECTORIES...]
```

### Options

- `-h, --help`: Show help message and exit
- `-v, --version`: Show version information and exit
- `-r, --recursive`: Recursively search subdirectories for files
- `-t, --type <extensions>`: File extensions to process (comma-separated, default: tf)
- `-e, --exclude <patterns>`: Exclude patterns (comma-separated) - directories and files to skip (from the current directory)
- `-`: Read HCL content from stdin

### Basic Usage

#### Validate files from stdin
```sh
cat config.tf | hcl-lint -
```

#### Validate specific files
```sh
hcl-lint main.tf variables.tf
```

#### Validate all .tf files in a directory (non-recursive)
```sh
hcl-lint /path/to/terraform/configs
```

#### Validate all .tf files in a directory and subdirectories
```sh
hcl-lint -r /path/to/terraform/configs
```

#### Validate .hcl files instead of .tf files
```sh
hcl-lint -t hcl ./config/
```

#### Validate multiple file types
```sh
hcl-lint -t tf,hcl,tfvars ./project/
```

#### Recursive validation with multiple file types
```sh
hcl-lint -r -t tf,hcl,tfvars ./project/
```

#### Validate multiple files and directories
```sh
hcl-lint main.tf ./modules ./environments
```

#### Exclude specific directories and files
```sh
# Exclude a directory
hcl-lint -r -e "tests/recursive" ./project/

# Exclude multiple patterns (directories and files)
hcl-lint -r -e "tests/recursive,.cache,dashboard/file.pp" ./project/

# Exclude with specific file types
hcl-lint -r -t tf,hcl -e ".git,node_modules,*.backup" ./project/
```
The exclude patterns are relative to the current directory.
If you want to exclude folder `./a/b/c`, you need to use `-e "a/b/c"`, not `-e "c"` or `-e "b/c"`.

## Examples

**Valid HCL file:**
```
$ hcl-lint -t hcl tests/valid.hcl
Checking tests/valid.hcl ... OK!
```

**Invalid HCL file:**
```
$ hcl-lint -t hcl tests/invalid.hcl
Checking tests/invalid.hcl ... Error parsing file: tests/invalid.hcl:13,1-2: Argument or block definition required; An argument or block definition is required here. To set an argument, use the equals sign "=" to introduce the argument value.
```

**Recursive directory validation:**
```
$ hcl-lint -r -t tf,hcl tests/
Checking tests/invalid.hcl ... Error parsing file: tests/invalid.hcl:13,1-2: Argument or block definition required; An argument or block definition is required here. To set an argument, use the equals sign "=" to introduce the argument value.
Checking tests/recursive/main.tf ... OK!
Checking tests/recursive/subdir1/subdir2/config.tf ... Error parsing file: tests/recursive/subdir1/subdir2/config.tf:9,33-34: Unclosed configuration block; There is no closing brace for this block before the end of the file.
Checking tests/recursive/subdir1/variables.hcl ... OK!
Checking tests/valid.hcl ... OK!
```

**No matching files:**
```
$ hcl-lint -t py tests/
No .py files found in directory 'tests/'
```

**Using exclude patterns:**
```
$ hcl-lint -r -t tf,hcl -e "tests/recursive" tests/
Checking tests/invalid.hcl ... Error parsing file: tests/invalid.hcl:13,1-2: Argument or block definition required; An argument or block definition is required here. To set an argument, use the equals sign "=" to introduce the argument value.
Checking tests/valid.hcl ... OK!
```

**Excluding multiple patterns:**
```
$ hcl-lint -r -e "tests/recursive,.cache,dashboard/file.pp" ./project/
# Only processes files not matching the exclude patterns
```

## Supported File Types

- `.tf` - Terraform configuration files (default)
- `.hcl` - HashiCorp Configuration Language files
- `.tfvars` - Terraform variable files
- Any file when explicitly specified (content will be parsed as HCL)

## Exit Codes

- `0`: All files are valid
- `1`: One or more files contain syntax errors

## Requirements

- Go 1.24 or later
- Dependencies are managed via Go modules