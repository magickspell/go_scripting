package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type productProviderMock struct {
	validSKU int64
}

func (m *productProviderMock) GetProduct(_ context.Context, sku int64) (*Product, error) {
	if sku == m.validSKU {
		return &Product{SKU: sku, Name: "Valid product"}, nil
	}
	return nil, ErrProductNotFound
}

func Test_httpHandler_getProduct(t *testing.T) {
	t.Run("invalid sku", func(t *testing.T) {
		// arrange
		productProviderM := &productProviderMock{}
		mux := productHandler(productProviderM)

		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/product/0", nil)

		mux.ServerHTTP(rec, req)

		if rec.Code != http.StatusBadRequest {
			t.Errorf("got %d, want %d", rec.Code, http.StatusBadRequest)
		}
	})

	t.Run("not found", func(t *testing.T) {
		productProviderM := &productProviderMock{}
		mux := productHandler(productProviderM)

		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/product/0", nil)

		mux.ServerHTTP(rec, req)

		if rec.Code != http.StatusNotFound {
			t.Errorf("got %d, want %d", rec.Code, http.StatusNotFound)
		}
	})

	t.Run("OK", func(t *testing.T) {
		productProviderM := &productProviderMock{validSKU: 1}
		mux := productHandler(productProviderM)

		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/product/1connectionSimple", nil)

		mux.ServerHTTP(rec, req)

		println(rec.Body.String())

		if rec.Code != http.StatusOK {
			t.Errorf("got %d, want %d", rec.Code, http.StatusOK)
		}
		if !strings.Contains(rec.Body.String(), `"sku":1connectionSimple`) {
			t.Errorf("unexpected body in response: %q", rec.Body.String())
		}
	})
}

func Test_helloHandler(t *testing.T) {
	t.Run("wrong method", func(t *testing.T) {
		mux := newRouter()
		// делаем моки для RESPONSE(ResponseRecorder) и REQUEST
		recTestMock := httptest.NewRecorder()
		reqTestMock := httptest.NewRequest(http.MethodPost, "/goodbye", nil)
		// делаем mock запрос
		mux.ServeHTTP(recTestMock, reqTestMock)
		if recTestMock.Code != http.StatusNotFound {
			t.Errorf("expected status code %v but got %v", http.StatusNotFound, recTestMock.Code)
		}
	})

	t.Run("ok", func(t *testing.T) {
		mux := newRouter()
		// делаем моки для RESPONSE(ResponseRecorder) и REQUEST
		recTestMock := httptest.NewRecorder()
		reqTestMock := httptest.NewRequest(http.MethodGet, "/hello", nil)
		// делаем mock запрос
		mux.ServeHTTP(recTestMock, reqTestMock)
		if recTestMock.Code != http.StatusOK {
			t.Errorf("expected status code %v but got %v", http.StatusOK, recTestMock.Code)
		}
	})
}
