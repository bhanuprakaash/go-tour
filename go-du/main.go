package main

import (
	"flag"
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
)

func main() {
	var root string
	flag.StringVar(&root, "path", ".", "Provide the path to analyze")

	var excludes Exclude
	flag.Var(&excludes, "exclude", "Provide the Exclude Fields")
	flag.Parse()

	excludeSet := make(map[string]struct{})
	for _, e := range excludes {
		excludeSet[strings.ToLower(e)] = struct{}{}
	}

	var analyzers = []Analyzer{
		&SizeStats{},
		&TypeStats{},
		&DuplicateStats{},
	}

	filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if _, found := excludeSet[strings.ToLower(info.Name())]; found {
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		if info.IsDir() {
			return nil
		}

		for _, analyzer := range analyzers {
			analyzer.Analyze(path, info)
		}

		return nil
	})

	for _, analyzer := range analyzers {
		switch v := analyzer.(type) {
		case *SizeStats:
			fmt.Printf("\n[DEBUG] SizeStats memory address: %p\n", v)
		case *TypeStats:
			fmt.Printf("\n[DEBUG] Tracking %d distinct file types\n", len(v.Extensions))
		default:
			fmt.Println("\n[DEBUG] Generic Analyzer")
		}

		analyzer.Report()
	}
}
