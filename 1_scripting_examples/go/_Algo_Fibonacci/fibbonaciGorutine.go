package main

import (
	"fmt"
)

// Функция для вычисления числа Фибоначчи с использованием горутин
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	// Каналы для получения результатов из горутин
	ch1 := make(chan int)
	ch2 := make(chan int)

	// Горутина для вычисления fibonacci(n-1)
	go func() {
		ch1 <- fibonacci(n - 1)
	}()

	// Горутина для вычисления fibonacci(n-2)
	go func() {
		ch2 <- fibonacci(n - 2)
	}()

	// Считывание результатов из каналов
	fib1 := <-ch1
	fib2 := <-ch2

	return fib1 + fib2
}

func main() {
	n := 30 // Например, вычислим 10-е число Фибоначчи
	fmt.Printf("Число Фибоначчи для %d: %d\n", n, fibonacci(n))
}
