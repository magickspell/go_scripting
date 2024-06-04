package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

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
