package main

import (
	"crypto/md5"
	"encoding/binary"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
)

const base62 = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

var (
	urlsMap = make(map[string]string)
	mu      sync.RWMutex
)

func hashUrl(url string) string {
	hash := md5.Sum([]byte(url))
	num := binary.BigEndian.Uint64(hash[:8])
	return encodeBase62(num)
}

func encodeBase62(n uint64) string {
	if n == 0 {
		return "0"
	}

	var out []byte
	for n > 0 {
		out = append(out, base62[n%62])
		n /= 62
	}

	// reverse
	for i, j := 0, len(out)-1; i < j; i, j = i+1, j-1 {
		out[i], out[j] = out[j], out[i]
	}

	return string(out)
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	code := strings.TrimPrefix(r.URL.Path, "/")
	log.Println("Requested code:", code)
	mu.RLock()
	originalURL, ok := urlsMap[code]
	mu.RUnlock()
	if !ok {
		http.NotFound(w, r)
		return
	}

	http.Redirect(w, r, originalURL, http.StatusFound)
}

func main() {
	url := "https://google.com"
	hashedUrl := hashUrl(url)

	mu.Lock()
	urlsMap[hashedUrl] = url
	mu.Unlock()

	fmt.Println(hashedUrl)
	fmt.Println("http://localhost:8080/" + hashedUrl)

	http.HandleFunc("/", redirectHandler)

	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
