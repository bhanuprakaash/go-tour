package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"io/fs"
	"os"
)

type DuplicateStats struct {
	FilesBySize map[int64][]string
}

func (d *DuplicateStats) Analyze(path string, info fs.FileInfo) {
	if d.FilesBySize == nil {
		d.FilesBySize = make(map[int64][]string)
	}
	if info.Size() < 1024 {
		return
	}
	d.FilesBySize[info.Size()] = append(d.FilesBySize[info.Size()], path)
}

func (d *DuplicateStats) Report(w io.Writer) {

	var wastedSpace int64
	var duplicateCount int

	for size, files := range d.FilesBySize {
		if len(files) < 2 {
			continue
		}

		hashMap := make(map[string][]string)

		for _, file := range files {
			hash, err := computeHash(file)
			if err != nil {
				continue
			}
			hashMap[hash] = append(hashMap[hash], file)
		}

		for _, sameFiles := range hashMap {
			if len(sameFiles) > 1 {
				duplicateCount++
				wasted := int64(len(sameFiles)-1) * size
				wastedSpace += wasted

				fmt.Fprintf(w, "Duplicate Group (%d bytes):\n", size)
				for _, f := range sameFiles {
					fmt.Fprintf(w, " - %s\n", f)
				}
			}
		}

	}

	if duplicateCount == 0 {
		fmt.Fprintln(w, "No duplicates found.")
	} else {
		fmt.Fprintf(w, "\nTotal Wasted Space: %.2f MB\n", float64(wastedSpace)/(1024*1024))
	}

}

func computeHash(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	h := sha256.New()

	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
