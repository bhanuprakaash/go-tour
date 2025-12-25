package main

import "fmt"

func main() {
	x := 1
	p := &x
	fmt.Println(*p)
	*p = 2
	fmt.Println(x)

	fmt.Println("--------local variable will remain in existence evenafter the call has returned-----------")
	var ptr = f()
	fmt.Println(*ptr)

	fmt.Println("---------increment----------")
	val := 1
	inc(&val)
	fmt.Println(val)

	fmt.Println("--------swap-----------")
	key1, key2 := 2, 3
	swap(&key1, &key2)
	fmt.Println("after swap: ", key1, key2)
}

func f() *int {
	v := 1
	return &v
}

func inc(p *int) int {
	*p++
	return *p
}

func swap(key1, key2 *int) {
	*key1, *key2 = *key2, *key1
}
