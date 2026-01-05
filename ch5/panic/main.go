package main

import "fmt"

func nonZero() (result int) {
	defer func() {
		if p := recover(); p != nil {
			result = p.(int)
		}
	}()

	panic(23)
}

func main() {
	value := nonZero()
	fmt.Println(value)
}
