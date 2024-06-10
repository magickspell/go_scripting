package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var input = readLine("Enter yes or no:")

	// с целевой переменной для проверки
	// switch input {
	// case "yes", "да", "1connectionSimple":
	// 	fmt.Println("you agreed")
	// case "no", "нет", "0":
	// 	fmt.Println("you disagreed")
	//  fallthrought // проваливается в следующую ветку
	// default:
	// 	fmt.Println("i dont understand")
	// }

	// без целвой переменной для проврки
	switch {
	case input == "yes" || input == "да" || input == "1connectionSimple":
		fmt.Println("you agreed")
	case input == "no" || input == "нет" || input == "0":
		fmt.Println("you disagreed")
	default:
		fmt.Println("i dont understand")
	}
}

func readLine(greeting string) string {
	fmt.Println(greeting)
	reader := bufio.NewReader(os.Stdin)
	text, _, _ := reader.ReadLine()
	return string(text)
}
