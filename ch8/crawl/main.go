package main

import (
	"fmt"
	"net/http"
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

	links = visit(links, n.FirstChild)
	links = visit(links, n.NextSibling)

	return links
}

func findLinks(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	return visit(nil, doc), nil
}

// func crawl1(url string) []string {
// 	fmt.Println(url)
// 	lists, err := findLinks(url)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	return lists
// }

var tokens = make(chan struct{}, 20)

func crawl2(url string) []string {
	fmt.Println(url)

	tokens <- struct{}{}
	lists, err := findLinks(url)
	<-tokens
	if err != nil {
		fmt.Println(err)
	}

	return lists
}

func main() {

	worklist := make(chan []string)
	unseenLinks := make(chan string)

	go func() {
		worklist <- os.Args[1:]
	}()

	for range 20 {
		go func() {
			for link := range unseenLinks {
				go func() {
					list := crawl2(link)
					worklist <- list
				}()
			}
		}()
	}

	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}

	close(worklist)

	fmt.Println("I'm Done.")

}
