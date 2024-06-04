package main

import (
	"log"
	"net/http"
)

type httpHandler struct {
	mux  *http.ServeMux
	name string
}

// создаем мультиплексор
func newHandler(name string) *httpHandler {
	handler := &httpHandler{
		mux:  http.NewServeMux(),
		name: name,
	}
	handler.mux.HandleFunc("/", handler.index)
	return handler
}

// просто перенаправляет запрос на мультиплексор
func (h *httpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mux.ServeHTTP(w, r)
}

// создаем метод для newHandler
func (h *httpHandler) index(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("index " + h.name))
}

func main() {
	mux := newHandler("jony")
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", mux))
}
