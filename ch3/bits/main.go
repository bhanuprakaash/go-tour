package main

import (
	"fmt"
	"math"
)

func main() {
	x := 3
	fmt.Printf("%08b", x)

	// for float  - use %f, %e, and %g

	for i := range 8 {
		fmt.Printf("%d, e = %8.3f\n", i, math.Exp(float64(i)))
	}

	var z float64

	fmt.Println(z, -z, 1/z, z/1, z/z, -1/z)

	var compx complex64 = complex(1, 2)
	var compy complex64 = complex(3, 4)

	fmt.Println(compx * compy)

}
