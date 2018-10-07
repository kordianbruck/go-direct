package src

import (
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Location", os.Getenv("URL"))
		w.WriteHeader(301)
	})

	if err := http.ListenAndServe(":80", nil); err != nil {
		panic(err)
	}
}
