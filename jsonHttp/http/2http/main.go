package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	httpClient := &http.Client{
		Transport:     &http.Transport{}, // механика выполнения запросов, RoundTripper интерфейс
		CheckRedirect: nil,               // функция определяющая политику редиректов
		Jar:           nil,               // общее хранилище куки
		Timeout:       3 * time.Second,   // лимит времени выполнения запроса
	}

	res, err := httpClient.Get("https://example.com")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("res.StatusCode", res.StatusCode)
}
