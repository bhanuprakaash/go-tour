package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

func worker(port string, wg *sync.WaitGroup) {
	defer wg.Done()
	address := fmt.Sprintf("localhost:%s", port)

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("server %s received request\n", port)
		fmt.Fprintf(w, "Hello from Server %s\n", port)
	})

	fmt.Printf("server listening on %s\n", port)
	err := http.ListenAndServe(address, mux)
	if err != nil {
		log.Fatal(err)
	}
}