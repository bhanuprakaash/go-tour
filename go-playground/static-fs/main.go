package main

import (
	"fmt"
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/index.html")
}

func FallbackHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/404.html")
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/{$}", homeHandler)

	mux.HandleFunc("/", FallbackHandler)

	fmt.Println("Starting server on :8080...")
	fmt.Println("Server is ready! Visit http://localhost:8080")

	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}
