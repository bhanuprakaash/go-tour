package main

import (
	"fmt"
	"os"
	"strings"
)

func sum(vals ...int) int {
	total := 0

	for _, x := range vals {
		total += x
	}
	return total
}

func errorf(linenum int, format string, args ...any) {
	fmt.Fprintf(os.Stderr, "Line %d: ", linenum)
	fmt.Fprintf(os.Stderr, format, args...)
	fmt.Fprintln(os.Stderr)
}

func max(vals ...int) int {
	if len(vals) == 0 {
		return 0
	}
	maxElement := vals[0]
	for _, item := range vals[1:] {
		if item > maxElement {
			maxElement = item
		}
	}
	return maxElement
}

func min(vals ...int) int {
	if len(vals) == 0 {
		return 0
	}
	minElement := vals[0]
	for _, item := range vals[1:] {
		if item < minElement {
			minElement = item
		}
	}
	return minElement
}

func joins(sep string, values ...string) string {
	if len(values) == 0 {
		return ""
	}

	n := 0
	for i, v := range values {
		n += len(v)
		if i > 0 {
			n += len(sep)
		}
	}

	var result strings.Builder
	result.Grow(n)
	for i, item := range values {
		if i > 0 {
			result.WriteString(sep)
		}
		result.WriteString(item)
	}

	return result.String()
}

func f(v ...int) {}
func g(v []int)  {}

func main() {
	values := []int{1, 2, 3, 4}
	// fmt.Println(sum(1, 2, 3, 4))
	fmt.Println(sum(values...))
	fmt.Printf("%T\n", f)
	fmt.Printf("%T\n", g)

	linenum, name := 12, "count"
	errorf(linenum, "undefined: %s", name)

	fmt.Println(max(1, 3, 4, 2323), min(-2, 3, 4, 2323))

	fmt.Println(joins(",", "I'm", "a", "very", "good", "guy"))
}
