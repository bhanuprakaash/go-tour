package main

import (
	"fmt"
	"time"
)

func main() {
	const x = 34

	const noDelay time.Duration = 0

	const timeout = 5 * time.Minute

	fmt.Printf("%T %[1]v\n", noDelay)
	fmt.Printf("%T\n", timeout)

	type Weekday int

	const (
		Sunday Weekday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)

	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)

	// 3.13
	const (
		KB = 1000
		MB = 1000 * KB
		GB = 1000 * MB
		TB = 1000 * GB
		PB = 1000 * TB
		EB = 1000 * PB
	)

	fmt.Println(KB, MB, GB, TB, PB, EB)

	const untypeConst = 92233720368547758079223372036854775807

	// var typeValue int64 = untypeConst

}
