package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Location", "https://google.com")
		w.WriteHeader(301)
	})

	if err := http.ListenAndServe(":80", nil); err != nil {
		panic(err)
	}
}
