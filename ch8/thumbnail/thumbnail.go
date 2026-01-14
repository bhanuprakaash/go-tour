package main

import (
	"fmt"
	"math/rand"
	"time"
)

type ThumbnailMock struct{}

func (t ThumbnailMock) ImageFile(infile string) (string, error) {
	delay := time.Duration(rand.Intn(500)) * time.Millisecond
	time.Sleep(delay)

	if rand.Float32() < 0.2 {
		return "", fmt.Errorf("failed to process image: %s", infile)
	}

	outfile := infile + ".thumb.jpg"
	return outfile, nil
}

func MockSize(filename string) int64 {
	return int64(rand.Intn(100000) + 1000)
}
