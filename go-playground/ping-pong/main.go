package main

import (
	"fmt"
	"sync"
	"time"
)

func player(table chan int, name string, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		ball, isOpen := <-table
		if !isOpen {
			fmt.Println("something went wrong")
			return
		}
		ball++
		if ball > 10 {
			fmt.Println("game over")
			return
		}

		fmt.Println(name, "hits", ball)
		time.Sleep(1 * time.Second)
		table <- ball
	}

}

func main() {
	var table = make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)

	go player(table, "ping", &wg)
	go player(table, "pong", &wg)

	table <- 1

	wg.Wait()

	fmt.Println("match finished")

}
