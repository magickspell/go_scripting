package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	str := "Пример строки"

	// обрабатывает каждый байт
	for i := 0; i < len(str); i++ {
		fmt.Printf("Символ на позиции %d: %c\n", i, str[i])
	}

	// брабатывает каждый символ
	for index, value := range str {
		fmt.Printf("Символ на позиции %d: %c\n", index, value)
	}

	// перегоняем строку в байты и обратно
	byteStr := "Hello, world"
	byteSlice := []byte(byteStr)
	runeSlice := []rune(byteStr)
	newStrByte := string(byteSlice)
	newStrRune := string(runeSlice)
	fmt.Println("byteStr:", byteStr)
	fmt.Println("byteSlice:", byteSlice)
	fmt.Println("runeSlice:", runeSlice)
	fmt.Println("newStrByte:", newStrByte)
	fmt.Println("newStrRune:", newStrRune)

	// склеивание строк
	// буфером
	var buffer bytes.Buffer
	buffer.WriteString("Hello, ")
	buffer.WriteString("world!")
	result := buffer.String()
	fmt.Println(result)
	// стринг билдером
	var builder strings.Builder
	builder.WriteString("Hello, ")
	builder.WriteString("world!")
	builderResult := builder.String()
	fmt.Println("builderResult:", builderResult)
}
