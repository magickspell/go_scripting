package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	helloHandler := func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "Hello, world!")
	}

	http.HandleFunc("/", helloHandler)

	handlerArgs := func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "%v\n", r)
		_, _ = fmt.Fprintln(w, "можно без форматирования")
		_, _ = w.Write([]byte("можно записать сразу в writer\n"))
	}
	http.HandleFunc("/handler-args", handlerArgs)

	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}
