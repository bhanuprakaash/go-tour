package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/bhanuprakaash/go-tour.git/ch2/tempconv"
)

func main() {
	inputs := os.Args[1:]

	for _, input := range inputs {
		val, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid number: %s\n", input)
			continue
		}
		fmt.Printf("%v = %v, %v = %v \n", tempconv.Celsius(val), tempconv.CToF(tempconv.Celsius(val)), tempconv.Fahrenheit(val), tempconv.FToC(tempconv.Fahrenheit(val)))
	}
}
