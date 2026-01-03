package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func countElements(n *html.Node, count map[string]int) {
	if n == nil {
		return
	}
	if n.Type == html.ElementNode {
		count[n.Data]++
	}

	countElements(n.FirstChild, count)
	countElements(n.NextSibling, count)

}

func main() {
	doc, err := html.Parse(os.Stdin)
	count := make(map[string]int)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}

	countElements(doc, count)

	for elementType, count := range count {
		fmt.Printf("%s:%d\n", elementType, count)
	}
}
