package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	// ages := make(map[string]int)

	// ages := map[string]int{
	// 	"alice": 34,
	// 	"mark":  24,
	// }

	ages := make(map[string]int)
	freq := make(map[string]int)
	text := "this is first text of this text"

	ages["alice"] = 23
	ages["mark"] = 12
	ages["antony"] = 32

	// show in sort
	names := make([]string, 0, len(ages))
	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}

	if age, ok := ages["bob"]; !ok {
		fmt.Printf("%T\n", age)
	}

	wordFreq(freq, text)

	fmt.Println("-------------------------")
	for word, count := range freq {
		fmt.Printf("%s\t%d\n", word, count)
	}
	fmt.Println("-------------------------")

}

func wordFreq(freq map[string]int, text string) {
	words := strings.Fields(text)
	for _, word := range words {
		freq[word]++
	}
}
