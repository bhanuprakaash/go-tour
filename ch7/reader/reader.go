package main

import (
	"fmt"
	"io"
	"os"
	// "os"
	// "time"
)

type MyReader struct {
	s      string
	offset int64
}

type Reader interface {
	Read(p []byte) (n int, err error)
}
type Closer interface {
	Close() error
}
type Writer interface {
	Write(p []byte) (n int, err error)
}
type ReadWriter interface {
	Reader
	Writer
}
type ReadWriteCloser interface {
	Reader
	Writer
	Closer
}

func (r *MyReader) Read(b []byte) (n int, e error) {
	if r.offset >= int64(len(r.s)) {
		return 0, io.EOF
	}
	n = copy(b, r.s[r.offset:])
	r.offset += int64(n)
	return n, nil
}

func NewReader(s string, offset int64) *MyReader {
	return &MyReader{s: s, offset: offset}
}

func main() {
	r := NewReader("hello world", 2)
	buf := make([]byte, 4)
	for {
		n, err := r.Read(buf)
		if n > 0 {
			fmt.Printf("Read %d bytes: %q\n", n, buf[:n])
		}
		if err == io.EOF {
			break
		}
	}

	// var w io.Writer
	// var rwc io.ReadWriteCloser
	// w = os.Stdout
	// // w = time.Second
	// w = rwc
	// rwc = w

	os.Stdout.Write([]byte("hello"))
	os.Stdout.Close()

	var byteStringArray = []byte{54, 23}
	var w io.Writer
	w = os.Stdout
	w.Write(byteStringArray)
	// w.Close()

}
