package main

import (
	"io"
	"io/fs"
)

type Analyzer interface {
	Analyze(path string, info fs.FileInfo)
	Report(w io.Writer)
}
