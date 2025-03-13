package main

import (
	"fmt"
	"time"
)

// Функция, которая будет выполняться в горутине
func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		// Приостанавливаем выполнение на 1 секунду
		time.Sleep(time.Second)
	}
}

func main() {
	// Запуск функции printNumbers в горутине
	go printNumbers()

	// Главная горутина также выполняет свою работу
	for i := 200; i <= 205; i++ {
		fmt.Println(i)
		// Приостанавливаем выполнение на 2 секунды
		time.Sleep(2 * time.Second)
	}
}
