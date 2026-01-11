package main

import "io/fs"

type Analyzer interface {
	Analyze(path string, info fs.FileInfo)
	Report()
}
