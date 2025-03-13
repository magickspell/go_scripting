package main

import "fmt"

func main() {
	// создать канал
	ch := make(chan int, 2)
	// записать данные в канал
	ch <- 1
	ch <- 2
	// ch <- 3 // panic
	fmt.Println(len(ch))
	fmt.Println(cap(ch))
	valueFromChannel := <-ch
	fmt.Println("valueFromChannel = ", valueFromChannel)
	fmt.Println(<-ch)
	ch <- 3
	fmt.Println(<-ch)
	close(ch) // закрываем канал, только отправитель закрывает канал
}
