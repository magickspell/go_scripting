package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("example.txt")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Файл не найден.")
		} else if os.IsPermission(err) {
			fmt.Println("Недостаточно прав для чтения файла.")
		} else {
			fmt.Println("Произошла неизвестная ошибка:", err)
		}
		return
	}
	defer file.Close()
	// Чтение файла
}
