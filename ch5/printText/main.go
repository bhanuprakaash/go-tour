package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func prinText(n *html.Node) {
	if n == nil {
		return
	}

	if n.Type == html.ElementNode &&
		(n.Data == "script" || n.Data == "style") {
		return
	}

	if n.Type == html.TextNode {
		text := strings.TrimSpace(n.Data)
		if text != "" {
			fmt.Println(text)
		}
	}

	prinText(n.FirstChild)
	prinText(n.NextSibling)
}

func main() {
	doc, err := html.Parse(os.Stdin)

	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}

	prinText(doc)
}
