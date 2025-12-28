package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var x string = "Hello, world"
	var y string = "Hello, world"
	// var str string = x[2:5]

	fmt.Println(x[2])

	fmt.Println(x[2:4])

	fmt.Println(x[:])

	fmt.Println(string(x[0]) < string(y[0]))

	fmt.Println(x + y + "3")

	fmt.Println("H" > "h")

	fmt.Printf("%p\n", unsafe.StringData(x))
	fmt.Printf("%p\n", unsafe.StringData(y))

	var lit string

	lit = `SELECT * FROM employees WHERE id=2;`

	fmt.Println(lit)

	for _, b := range []byte(string('❌')) {
		fmt.Printf("%08b ", b)
	}

	s := []byte(x)
	fmt.Println(s)

}
