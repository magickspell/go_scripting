package main

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var ErrProductNotFound = errors.New("product not found")

type httpHandler struct {
	router          *mux.Router
	productProvider ProductProvider
}

type Product struct {
	SKU  int64  `json:"sku"`
	Name string `json:"name"`
}

// ProductProvider отдает товар по sku
type ProductProvider interface {
	GetProduct(ctx context.Context, sku int64) (*Product, error)
}

func productHandler(productProvider ProductProvider) *httpHandler {
	router := mux.NewRouter()
	handler := &httpHandler{
		router:          router,
		productProvider: productProvider,
	}
	router.HandleFunc("/product/{sku:[0-9]+}", handler.getProduct).Methods(http.MethodPost)
	return handler
}

func (h *httpHandler) ServerHTTP(w http.ResponseWriter, req *http.Request) {
	h.router.ServeHTTP(w, req)
}

func (h *httpHandler) getProduct(w http.ResponseWriter, req *http.Request) {
	// получаем SKU товара из запроса
	vars := mux.Vars(req)
	sku, _ := strconv.ParseInt(vars["sku"], 10, 64)
	if sku == 0 {
		println("### ZERO FOUND ###")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product, err := h.productProvider.GetProduct(req.Context(), sku)
	if errors.Is(err, ErrProductNotFound) { // так правильнее ропверять ошибки
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	result, err := json.Marshal(product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	_, _ = w.Write(result)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("hello world"))
}

func newRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/hello", helloHandler).Methods(http.MethodGet)
	//router.HandleFunc("/get-product", productHandler).Methods(http.MethodPost)
	return router
}
