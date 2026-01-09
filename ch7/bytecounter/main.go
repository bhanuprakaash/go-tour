package main

import (
	"bufio"
	"fmt"
	"strings"
)

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(string(p)))
	// scanner.Split(bufio.ScanWords)
	scanner.Split(bufio.ScanLines)
	count := 0

	for scanner.Scan() {
		count++
	}

	*c += ByteCounter(count)
	return len(p), nil
}

func main() {
	var c ByteCounter
	c.Write([]byte("hello \n world"))
	fmt.Println(c)

	c = 0
	var name = "Dolly"
	fmt.Fprintf(&c, "Hello, %s", name)
	fmt.Println(c)
}
