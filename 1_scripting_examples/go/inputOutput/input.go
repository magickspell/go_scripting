package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// считывает данные разделяи пробелами
	var name string
	fmt.Print("Введите ваше имя: ")
	fmt.Scan(&name)
	fmt.Printf("Привет, %s!\n", name)

	// буфер более гибкий,можно читать строки - как будто самое нормальное
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите ваше сообщение: ")
	message, _ := reader.ReadString('\n')
	fmt.Printf("Вы ввели: %s", message)

	// os.Stdin более низкий уровень чтения
	fmt.Print("Введите символы: ")
	buffer := make([]byte, 10) // Буфер на 10 байт
	n, err := os.Stdin.Read(buffer)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Считано %d байт: %s\n", n, string(buffer))
}
