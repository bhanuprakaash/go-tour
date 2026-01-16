package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	done := make(chan struct{})
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	var n int

	go func() {
		os.Stdin.Read(make([]byte, 1))
		reader := bufio.NewReader(os.Stdin)
		_, _ = reader.ReadBytes('\n')
		close(done)
	}()

	fmt.Println("countdown starts...")

	for {
		select {
		case <-ticker.C:
			n++
			fmt.Println(n)
		case <-done:
			fmt.Println("clean up done")
			return
		}
	}

}
