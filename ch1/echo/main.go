package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	//1.1
	fmt.Println(strings.Join(os.Args[0:], " "))

	// //1.2
	for i, arg := range os.Args[1:] {
		fmt.Printf("%d %s \n", i, arg)
	}

	//1.3
	//  go run main.go $(seq 1 10000)
	start := time.Now()
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Printf("%.7fs elapse using loops\n", time.Since(start).Seconds())

	joinStart := time.Now()
	strings.Join(os.Args[0:], " ")
	fmt.Printf("%.7fs elapse using joins\n", time.Since(joinStart).Seconds())

}
