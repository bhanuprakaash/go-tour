package main

import (
	"fmt"
	"io"
	"os"
)

type wrapper struct {
	originalWriter io.Writer
	count          int64
}

func (w *wrapper) Write(p []byte) (int, error) {
	n, err := w.originalWriter.Write(p)
	w.count += int64(n)

	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	newWrapper := &wrapper{
		originalWriter: w,
		count:          0,
	}

	return newWrapper, &newWrapper.count
}

func main() {
	newWriter, counterPtr := CountingWriter(os.Stdout)
	fmt.Fprintf(newWriter, "Hello\n")
	fmt.Println("Bytes written so far:", *counterPtr) // Output: 6
}
