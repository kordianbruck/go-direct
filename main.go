package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Location", os.Getenv("URL"))
	})
	http.ListenAndServe(":80", nil)
}
