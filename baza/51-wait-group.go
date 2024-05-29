package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(1)

	var wg sync.WaitGroup // объявляет переменную wg типа sync.WaitGroup. Этот тип используется для ожидания завершения всех горутин.
	for x := 0; x < 10; x++ {
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			fmt.Printf("gorutine #{x}\n")
		}(x)
	}
}
