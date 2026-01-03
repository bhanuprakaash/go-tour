package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func visit(links []string, n *html.Node) []string {

	if n == nil {
		return links
	}

	if n.Type == html.ElementNode && n.Data == "a" {
		for _, attr := range n.Attr {
			if attr.Key == "href" {
				links = append(links, attr.Val)
			}
		}
	}

	if n.Type == html.ElementNode && n.Data == "img" {
		for _, attr := range n.Attr {
			if attr.Key == "src" {
				links = append(links, attr.Val)
			}
		}
	}

	if n.Type == html.ElementNode && n.Data == "link" {
		var isStyleSheet bool
		var href string
		for _, attr := range n.Attr {
			if attr.Key == "rel" && attr.Val == "stylesheet" {
				isStyleSheet = true
			}
			if attr.Key == "href" {
				href = attr.Val
			}
		}
		if isStyleSheet && href != "" {
			links = append(links, href)
		}
	}

	if n.Type == html.ElementNode && n.Data == "script" {
		for _, attr := range n.Attr {
			if attr.Key == "src" {
				links = append(links, attr.Val)
			}
		}
	}

	// for c := n.FirstChild; c != nil; c = c.NextSibling {
	// 	links = visit(links, c)
	// }

	links = visit(links, n.FirstChild)
	links = visit(links, n.NextSibling)

	return links
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}
