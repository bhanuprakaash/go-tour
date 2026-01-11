package main

import (
	"flag"
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
		analyzer.Report()
	}
}
