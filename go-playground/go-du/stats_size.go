package main

import (
	"fmt"
	"io"
	"io/fs"
)

type SizeStats struct {
	TotalSize int64
	FileCount int
}

func (s *SizeStats) Analyze(path string, info fs.FileInfo) {
	s.FileCount++
	s.TotalSize += info.Size()
}

func (s *SizeStats) Report(w io.Writer) {
	fmt.Fprintf(w, "Files Found: %d\n", s.FileCount)
	fmt.Fprintf(w, "Total Files Size: %.2f\n", float64(s.TotalSize)/(1024*1024))
}

func (s *SizeStats) Reset() {
	s.TotalSize = 0
}
