package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	urls := []string{
		"https://www.lamoda.ru",
		"https://www.ya.ru",
		"https://www.mail.ru",
		"https://www.google.ru",
	}

	waitGroup := &sync.WaitGroup{}

	for _, url := range urls {
		waitGroup.Add(1)
		go func(url string) {
			defer waitGroup.Done()

			fmt.Printf("Fetching: %v...\n", url)

			err := fetchUrl(url)
			if err != nil {
				fmt.Printf("Error fetching: %s \n", url)
				fmt.Println(err)
				return
			}
		}(url)
	}

	fmt.Println("All request launched!")
	// time.Sleep(400 * time.Millisecond) // ждем 400 милисекунд
	waitGroup.Wait() // ждем пока все зафетчится
	fmt.Println("Program finished")
}

func fetchUrl(url string) error {
	_, err := http.Get(url)
	return err
}
