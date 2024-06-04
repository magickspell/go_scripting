package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("hello world"))
}

func newRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/hello", helloHandler).Methods(http.MethodGet)
	return router
}
