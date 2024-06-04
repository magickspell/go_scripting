package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

const (
	url         = "https://httpbin.org/post"
	contentType = "application/json"
	reqBody     = `{"id": 999, "value": "content"}`
)

func main() {
	httpClient := &http.Client{Timeout: time.Second * 10}

	{
		req, _ := http.NewRequest(http.MethodPost, url, nil)
		req.Header.Set("Content-Type", contentType)

		resp, err := httpClient.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body)
		fmt.Println("body", string(body))
	}

	{
		// No context, only POST (СОКРАЩЕННО)
		// способ не особо хороший т.к. нельзя передать контекст запроса и можно только ГЕТ и ПОСТ
		resp, err := httpClient.Post(url, contentType, strings.NewReader(reqBody))
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body)
		fmt.Println("body", string(body))
	}
}
