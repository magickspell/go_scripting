package main

import (
	"fmt"
	"sync"
)

var counter int = 0

func increment() {
	counter++
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			increment()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Значение счетчика:", counter)
}
