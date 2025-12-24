package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]map[string]bool)
	files := os.Args[1:]

	if len(files) == 0 {
		err := fmt.Errorf("No file found")
		fmt.Fprintln(os.Stderr, err)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, filenames := range counts {
		if len(filenames) > 1 {
			fmt.Printf("Line: '%s'\nFound in: ", line)
			for fname := range filenames {
				fmt.Printf("%s", fname)
			}
			fmt.Println()
		}
	}
}

func countLines(f *os.File, counts map[string]map[string]bool) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		text := input.Text()
		if counts[text] == nil {
			counts[text] = make(map[string]bool)
		}

		counts[text][f.Name()] = true
	}
}
