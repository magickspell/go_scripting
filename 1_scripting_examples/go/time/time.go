package main

import (
	"fmt"
	"time"
)

func main() {
	// Получение текущего времени
	now := time.Now()
	fmt.Println("Current time:", now)

	// Извлечение компонентов времени
	fmt.Println("Hour:", now.Hour())
	fmt.Println("Minute:", now.Minute())
	fmt.Println("Second:", now.Second())

	// Форматирование времени в строку
	fmt.Println("Formatted time:", now.Format("2006-01-02 15:04:05"))

	// Создание временного интервала
	duration := time.Hour * 24 // 24 часа
	fmt.Println("Duration:", duration)

	// Выполнение арифметических операций с временем
	future := now.Add(duration)
	fmt.Println("Future time:", future)
}
