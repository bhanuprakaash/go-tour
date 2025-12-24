package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"time"
)

func main() {
	start := time.Now()

	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	defer resp.Body.Close()
	filename := path.Base(url) + ".txt"
	destFile, err := os.Create(filename)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	defer destFile.Close()
	_, err = io.Copy(destFile, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("while reading%s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %s", secs, url)
}
