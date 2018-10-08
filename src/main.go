package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	url, ok := os.LookupEnv("URL")
	if !ok {
		url = "https://google.com"
	}
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "80"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Location", url)
		w.WriteHeader(301)
	})

	fmt.Printf("Starting redirection server on port %v with redirection to %v", port, url)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		panic(err)
	}
}
