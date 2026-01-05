package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()
	local := path.Base(resp.Request.URL.Path)
	fmt.Println(local)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	n, err = io.Copy(f, resp.Body)

	defer func() {
		closeErr := f.Close()
		if err == nil {
			err = closeErr
		}
	}()
	return local, n, err
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <url>\n", os.Args[0])
		os.Exit(1)
	}

	// 2. Get the URL from the arguments
	url := os.Args[1]

	// 3. Call the fetch function
	filename, n, err := fetch(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fetch failed: %v\n", err)
		os.Exit(1)
	}

	// 4. Print success message
	fmt.Printf("Successfully downloaded %s (%d bytes)\n", filename, n)
}
