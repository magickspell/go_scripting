package main

import (
	"fmt"
	"os"
)

func main() {
	// Создаем файл. Если файл уже существует, он будет открыт для записи.
	file, err := os.Create("example.txt")
	if err != nil {
		// Обработка возможной ошибки при создании файла
		fmt.Println("Ошибка при создании файла:", err)
		return
	}
	defer file.Close() // Закрытие файла при выходе из функции

	// Данные для записи
	data := "Привет, мир!"

	// Запись данных в файл
	_, err = file.WriteString(data)
	if err != nil {
		// Обработка ошибки записи
		fmt.Println("Ошибка при записи в файл:", err)
		return
	}

	fmt.Println("Файл успешно записан.")
}
