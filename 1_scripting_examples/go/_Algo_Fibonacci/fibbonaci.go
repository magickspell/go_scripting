package main

import "fmt"

// Функция для вычисления числа Фибоначчи
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	// Пример использования функции
	n := 40 // Например, вычислим 10-е число Фибоначчи
	fmt.Printf("Число Фибоначчи для %d: %d\n", n, fibonacci(n))
}
