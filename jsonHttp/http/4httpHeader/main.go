package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func main() {
	fmt.Println()

	req, err := http.NewRequest(http.MethodGet, "https://reqres.in/api/users", nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("Accepts", "application/json")
	req.Header.Set("CusTom", "CusTom")

	q := url.Values{} // map[string][]string
	q.Add("page", "2")
	q.Add("paGE", "3") // case sensetive

	req.URL.RawQuery = q.Encode()

	fmt.Println()
	fmt.Println("header", req.Header)
	fmt.Println()
	fmt.Println("URL", req.URL)
}
