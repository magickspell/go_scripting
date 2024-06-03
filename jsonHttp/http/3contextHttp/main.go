package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

// эталонный вариант отправки запросов (лучший вариант)
func main() {
	fmt.Println()

	httpClient := &http.Client{Timeout: time.Second * 5}

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "https://example.com", nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	fmt.Println("statusCode", resp.StatusCode)
}
