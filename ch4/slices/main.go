package main

import (
	"fmt"
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
