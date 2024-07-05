# Public Suffix List Filter (psl)

This Go program fetches and filters the Public Suffix List based on user-specified criteria.

## Features

- Fetches the latest Public Suffix List from https://publicsuffix.org/
- Filters the list based on either a comment substring or a domain substring
- Command-line interface for easy use

## Prerequisites

- Go 1.19 or higher

## Installation

Assuming you have [Go](https://go.dev/doc/install) installed:

```bash
go install github.com/cybercdh/psl@latest
```

## Usage

Run the program using the following command:

```
psl [options]
```

### Options

- `-c string`: Filter by comment substring
- `-d string`: Filter by domain substring

You must provide either the `-c` or `-d` option.

### Examples

Filter by comment:
```bash
psl -c "ICANN"
```

Filter by domain:
```bash
psl -d "amazonaws.com"
```

## How it works

1. The program fetches the Public Suffix List from the specified URL.
2. It then filters the list based on the provided criteria:
   - If a comment substring is provided, it searches for the substring in comment lines and captures the subsequent non-comment lines until a blank line is encountered.
   - If a domain substring is provided, it filters non-comment lines that contain the specified substring.
3. The filtered results are printed to the console.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

[MIT](https://choosealicense.com/licenses/mit/)