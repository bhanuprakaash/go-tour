package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("Starting the countdown")
	tick := time.Tick(1 * time.Second)
	abort := make(chan struct{})

	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()
	for i := 10; i > 0; i-- {
		fmt.Println(i)
		select {
		case <-tick:
			// do nothing
		case <-abort:
			fmt.Println("aborting...")
			return
		}
	}
}
