package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type getQoutContract struct {
	Quote string `json:"quote"`
}

func getQuote(ctx context.Context) (*getQoutContract, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*30)
	defer cancel()

	req, _ := http.NewRequest("GET", "https://api.kanye.rest/", nil)

	httpClient := &http.Client{Timeout: time.Second * 10}

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("httpClinbt.Do: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body) // считать все из боди (там больше ниего не останется - ридер будет пустой)
	if err != nil {
		return nil, fmt.Errorf("ioutil.ReadAll: %w", err)
	}

	contract := new(getQoutContract)
	err = json.Unmarshal(respBody, contract)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %w", err)
	}

	return contract, nil
}

func main() {
	fmt.Println()

	ctx := context.Background()
	contract, err := getQuote(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Kane West says %q\n", contract.Quote)
}
