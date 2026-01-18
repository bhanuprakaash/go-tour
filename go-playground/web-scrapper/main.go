package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"golang.org/x/net/html"
)

type Result struct {
	url   string
	title string
	err   error
}

func getTitle(url string) (string, error) {
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	tokenizer := html.NewTokenizer(resp.Body)

	for {
		tokenType := tokenizer.Next()
		switch tokenType {
		case html.ErrorToken:
			return "", fmt.Errorf("title not found")
		case html.StartTagToken:
			token := tokenizer.Token()
			if token.Data == "title" {
				tokenizer.Next()
				return tokenizer.Token().Data, nil
			}
		}
	}
}

func worker(urls <-chan string, wg *sync.WaitGroup, results chan<- Result, ticker *time.Ticker) {
	defer wg.Done()
	for url := range urls {
		<-ticker.C
		title, err := getTitle(url)
		results <- Result{
			url:   url,
			title: title,
			err:   err,
		}
	}

}

func main() {
	urls := []string{
		"https://google.com",
		"https://bhanuprakashsai.com",
		"https://pkg.go.dev/net/http",
		"https://go.dev",
		"https://o.dev",
	}
	var wg sync.WaitGroup
	var resultsChan = make(chan Result, len(urls))
	var workers = 3
	var urlsChan = make(chan string)
	rate := time.Second / 2
	ticker := time.NewTicker(rate)
	defer ticker.Stop()

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go worker(urlsChan, &wg, resultsChan, ticker)
	}

	for _, url := range urls {
		urlsChan <- url
	}
	close(urlsChan)

	go func() {
		wg.Wait()
		close(resultsChan)
	}()

	for res := range resultsChan {
		if res.err != nil {
			fmt.Printf("[Error] %s: %v\n", res.url, res.err)
		} else {
			fmt.Printf("[Success] %s -> %s\n", res.url, res.title)
		}
	}

	fmt.Println("[Done]")

}
