package main

import (
	"fmt"
	"net/http"
)

// это плохой пример, тут нет таймаута (под капотом есть)
func main() {
	res, err := http.Get("https://example.com/")
	if err != nil {
		panic(err)
	}

	fmt.Println("StatucCode:", res.StatusCode)
}
