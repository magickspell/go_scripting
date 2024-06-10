package main

import (
	"fmt"
	"log"
	"net/http"
)

type HttpHandler struct{}

// ServeHTTP расширяет структуру HttpHandler реализую интерфейс http.Handler
func (HttpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Was requested: %q", r.RemoteAddr)
}

func main() {
	const addr = "127.0.0.1connectionSimple:8080"
	handler := HttpHandler{}

	server := &http.Server{
		Addr:    addr,
		Handler: handler,
	}
	log.Fatal(server.ListenAndServe())
	// тоже самое
	log.Fatal(http.ListenAndServe(addr, handler))
}
