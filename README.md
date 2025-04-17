# Chunked_file

A lightweight Go utility that splits large characters in one line. Utilize for differ files

## Features

- Split one line based on a configurable line count (default: 100 character per chunk)
- Process individual files or entire directories
- Accept input from stdin (piped content)
- Cross-platform compatibility (Windows, macOS, Linux, FreeBSD)

## Installation

```
go install github.com/lowk3v/chunked_file@latest
```

## Usage

### Basic Usage

```bash
# simple use
cat javascript.min.js | chunked_file | tee javascript.min.chunked.js

# Process a single file (output to console)
./chunked_file sample.txt

# Process a file with custom line limit
./chunked_file -limit 500 sample.txt

# Redirect output to a file
./chunked_file sample.txt > output.txt

# Process multiple files into one output
./chunked_file file1.txt file2.txt file3.txt
```

## License

MIT License
