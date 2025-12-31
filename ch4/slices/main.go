package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	months := []string{
		1:  "Jan",
		2:  "Feb",
		3:  "Mar",
		4:  "Apr",
		5:  "May",
		6:  "Jun",
		7:  "Jul",
		8:  "Aug",
		9:  "Sep",
		10: "Oct",
		11: "Nov",
		12: "Dec",
	}
	q := make([]int, 3, 7)

	array := [4]int{1, 2, 3, 4}
	adjacent := []int{1, 2, 2, 3, 4}

	Q2 := months[4:7]
	summer := months[6:9]
	fmt.Println(len(Q2), cap(Q2))
	fmt.Println(len(summer), cap(summer))

	q = append(q, 23)
	q = append(q, 27, 28, 29)

	fmt.Println(q, cap(q))

	q = append(q, 100)

	fmt.Println(q, cap(q))

	k := q[1:]

	k[1] = 99

	fmt.Println(k)
	fmt.Println(q)

	data := []string{"one", "", "three"}
	data = nonempty(data)

	fmt.Println(data, len(data), cap(data))

	data = nonempty2(data)
	fmt.Println(data, len(data), cap(data))

	s := []int{5, 6, 7, 8, 9}
	fmt.Println(remove(s, 2))
	reverse(&array)
	fmt.Println(array)

	q = rotate(q, 2)

	fmt.Println(q)

	adjacent = compareAdjacent(adjacent)

	fmt.Println(adjacent)

	bb := []byte{}

	bb = append(bb, 65, 65)

	fmt.Println(string(bb))

	str := "Hello 世界"

	// Convert to mutable byte slice
	b := []byte(str)

	fmt.Printf("Original: %s\n", b)

	reverseUtf8(b)
	fmt.Printf("Reversed: %s\n", b)

}

func nonempty(strings []string) []string {
	i := 0

	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}

	return strings[:i]
}

func nonempty2(strings []string) []string {
	out := strings[:0]

	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func reverse(array *[4]int) {
	for i, j := 0, len(array)-1; i < j; i, j = i+1, j-1 {
		array[i], array[j] = array[j], array[i]
	}
}

func rotate(s []int, k int) []int {
	if len(s) == 0 {
		return s
	}
	k = k % len(s)
	return append(s[k:], s[:k]...)
}

func compareAdjacent(s []int) []int {
	if len(s) == 0 {
		return s
	}

	writer := 1

	for i := 1; i < len(s); i++ {
		if s[i-1] != s[i] {
			s[writer] = s[i]
			writer++
		}
	}

	return s[:writer]
}

func reverseBytes(b []byte) {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
}
func reverseUtf8(b []byte) {
	reverseBytes(b)

	start := 0
	for i := range b {
		if utf8.RuneStart(b[i]) {
			reverseBytes(b[start:i])
		}
		start = i + 1
	}
}
