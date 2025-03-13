package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// конвертируем В JSON строку
	person := Person{Name: "Alice", Age: 30}
	jsonBytes, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Ошибка при маршалировании:", err)
		return
	}
	fmt.Println("[jsonBytes]")
	fmt.Println(string(jsonBytes))

	// конвертируем ИЗ JSON строку
	jsonStr := `{"name":"Bob","age":25}`
	var newPersonFromJson Person
	if err := json.Unmarshal([]byte(jsonStr), &newPersonFromJson); err != nil {
		fmt.Println("Ошибка при демаршалировании:", err)
		return
	}
	fmt.Println("[jsonBytes]")
	fmt.Println("Имя:", newPersonFromJson.Name)
	fmt.Println("Возраст:", newPersonFromJson.Age)
}
