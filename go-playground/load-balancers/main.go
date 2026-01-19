package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
)

var (
	serverPool ServerPool
	wg         sync.WaitGroup
)

func lbHandler(w http.ResponseWriter, r *http.Request) {
	targetUrl := serverPool.GetNextPeer()

	proxy := httputil.NewSingleHostReverseProxy(targetUrl)

	r.URL.Host = targetUrl.Host
	r.URL.Scheme = targetUrl.Scheme
	r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))
	r.Host = targetUrl.Host

	proxy.ServeHTTP(w, r)

}

func main() {
	ports := []string{"8081", "8082", "8083"}

	for _, port := range ports {
		wg.Add(1)
		go worker(port, &wg)
	}

	for _, port := range ports {
		address := fmt.Sprintf("http://localhost:%s", port)
		u, _ := url.Parse(address)
		serverPool.AddToBackend(u)
	}

	server := http.Server{
		Addr:    ":8000",
		Handler: http.HandlerFunc(lbHandler),
	}
	fmt.Printf("Load Balancer started at :8000\n")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

	wg.Wait()
	fmt.Println("[done]")

}
