package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Define command line flags
	limit := flag.Int("limit", 100, "Maximum number of lines per chunk")
	flag.Parse()

	// Get files from arguments
	args := flag.Args()

	// Process files from arguments or stdin
	if len(args) == 0 {
		// No arguments, process stdin
		processReader(os.Stdin, *limit)
	} else {
		// Process each argument (file or directory)
		for _, arg := range args {
			info, err := os.Stat(arg)
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "Error accessing %s: %v\n", arg, err)
				continue
			}

			if info.IsDir() {
				// Process directory
				processDirectory(arg, *limit)
			} else {
				// Process single file
				processFile(arg, *limit)
			}
		}
	}
}

// processDirectory walks through the directory and processes each file
func processDirectory(dirPath string, limit int) {
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Error accessing path %s: %v\n", path, err)
			return nil // continue walking
		}

		if !info.IsDir() {
			processFile(path, limit)
		}
		return nil
	})

	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error walking directory %s: %v\n", dirPath, err)
	}
}

// processFile opens and processes a single file
func processFile(filePath string, limit int) {
	file, err := os.Open(filePath)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error opening file %s: %v\n", filePath, err)
		return
	}
	defer file.Close()

	_, _ = fmt.Fprintf(os.Stderr, "Processing file: %s\n", filePath)
	processReader(file, limit)
}

// processReader reads from a reader and processes content in chunks
func processReader(reader io.Reader, limit int) {
	scanner := bufio.NewReader(reader)

	charCount := 0
	chunk := make([]string, 0, limit)

	// read char by char
	for {
		char, _, err := scanner.ReadRune()
		if err == io.EOF || err != nil {
			break
		}
		chunk = append(chunk, string(char))
		charCount++

		// When we reach the limit, process the chunk
		if charCount >= limit {
			fmt.Println(strings.Join(chunk, ""))
			charCount = 0
			chunk = make([]string, 0, limit)
		}
	}

	// Process any remaining lines
	if charCount > 0 {
		fmt.Println(strings.Join(chunk, ""))
	}
}
