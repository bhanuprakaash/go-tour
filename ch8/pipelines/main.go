package main

import "fmt"

func counter(out chan<- int) {
	for i := range 100 {
		out <- i
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for x := range in {
		out <- x * x
	}
	close(out)
}

func printers(in <-chan int) {
	for x := range in {
		fmt.Println(x)
	}
}

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go counter(naturals)
	go squarer(squares, naturals)
	printers(squares)

	// go func() {
	// 	for i := range 100 {
	// 		naturals <- i
	// 	}
	// 	close(naturals)
	// }()

	// go func() {
	// 	for x := range naturals {
	// 		squares <- x * x
	// 	}
	// 	close(squares)
	// }()

	// for x := range squares {
	// 	fmt.Println(x)
	// }
}
