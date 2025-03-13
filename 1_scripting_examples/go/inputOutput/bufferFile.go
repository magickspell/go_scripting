package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("buffered.txt")
	if err != nil {
		fmt.Println("Ошибка при создании файла:", err)
		return
	}
	defer file.Close()

	// Создаем буферизированный писатель
	writer := bufio.NewWriter(file)

	// Записываем данные
	_, err = writer.WriteString("Это строка будет записана с использованием буферизации\n")
	if err != nil {
		fmt.Println("Ошибка при записи данных:", err)
		return
	}

	// Важно не забыть сбросить данные из буфера в файл
	err = writer.Flush()
	if err != nil {
		fmt.Println("Ошибка при сбросе данных:", err)
	}

	// ЧИТАЕЕМ ФАЙЛ
	fileReade, err := os.Open("buffered.txt")
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer fileReade.Close()

	// Создаем буферизированный читатель
	reader := bufio.NewReader(fileReade)

	// Читаем данные
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка при чтении данных:", err)
		return
	}

	fmt.Print("Прочитанная строка: ", line)
}
