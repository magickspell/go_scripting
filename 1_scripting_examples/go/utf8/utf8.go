package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "Пример текста"

	// Получение длины строки в байтах
	length := len(str)
	fmt.Printf("Длина строки в байтах: %d\n", length)

	// Получение длины строки в символах Unicode
	runeCount := utf8.RuneCountInString(str)
	fmt.Printf("Длина строки в символах: %d\n", runeCount)

	// Итерирование по строке и вывод отдельных символов
	for i, r := range str {
		fmt.Printf("Символ %d: %c\n", i, r)
	}
}
