package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		words, images, err := CountWordsAndImages(url)

		if err != nil {
			fmt.Fprintf(os.Stderr, "count error: %v\n", err)
		}

		fmt.Printf("images: %d\nwords: %d\n", images, words)
	}
}

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImagesFromNode(doc)
	return
}

func countWordsAndImagesFromNode(n *html.Node) (words, images int) {
	if n == nil {
		return 0, 0
	}

	if n.Type == html.ElementNode && n.Data == "img" {
		images++
	}

	if n.Type == html.TextNode && n.Parent != nil {
		parent := n.Parent.Data
		if parent != "script" && parent != "style" {
			words += len(strings.Fields(n.Data))
		}
	}

	w1, i1 := countWordsAndImagesFromNode(n.FirstChild)
	w2, i2 := countWordsAndImagesFromNode(n.NextSibling)

	return words + w1 + w2, images + i1 + i2
}
