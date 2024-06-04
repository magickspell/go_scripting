package main

import (
	"fmt"
	"log"
	"net/http"
)

// кароче мультиплексоры не удобные и обычно делают свои кастомные см 4 папку

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/1", func(w http.ResponseWriter, req *http.Request) {
		_, _ = fmt.Fprintln(w, "one")
	})

	mux.HandleFunc("/2", func(w http.ResponseWriter, req *http.Request) {
		_, _ = fmt.Fprintln(w, "two")
	})

	log.Fatal(http.ListenAndServe("127.0.0.1:8080", mux))
}
