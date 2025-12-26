package main

import "fmt"

type Celsius float64
type Fahrenheit float64

type Miles float64
type Kilometers float64

func main() {
	var c Celsius = 100
	var f Fahrenheit = Fahrenheit(c)

	typeDiff(c)

	fmt.Println(c)
	fmt.Println(c == Celsius(f))
	fmt.Println(c == 22)
	fmt.Println(f - Fahrenheit(c))
}

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func typeDiff(a Celsius) Celsius { return a }

func MtoK(m Miles) Kilometers {
	return Kilometers(m * 1.609)
}
