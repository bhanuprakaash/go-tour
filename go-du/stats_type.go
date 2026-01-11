package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"sort"
)

type TypeCount struct {
	Ext   string
	Count int
}
type ByCount []TypeCount

func (b ByCount) Len() int           { return len(b) }
func (b ByCount) Less(i, j int) bool { return b[i].Count < b[j].Count }
func (b ByCount) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }

type TypeStats struct {
	Extensions map[string]int
}

func (t *TypeStats) Analyze(path string, info fs.FileInfo) {
	if t.Extensions == nil {
		t.Extensions = make(map[string]int)
	}
	ext := filepath.Ext(path)
	t.Extensions[ext]++
}

func (t *TypeStats) Report() {
	var stats []TypeCount
	for ext, count := range t.Extensions {
		stats = append(stats, TypeCount{Ext: ext, Count: count})
	}

	sort.Sort(sort.Reverse(ByCount(stats)))

	for _, s := range stats {
		fmt.Printf("%-10s: %d\n", s.Ext, s.Count)
	}
}
