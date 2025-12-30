package main

import "fmt"

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

}
