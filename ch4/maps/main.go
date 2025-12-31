package main

import (
	"fmt"
	"sort"
)

func main() {
	// ages := make(map[string]int)

	// ages := map[string]int{
	// 	"alice": 34,
	// 	"mark":  24,
	// }

	ages := make(map[string]int)

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
		fmt.Printf("%T", age)
	}


}

