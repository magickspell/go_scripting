package main

import "fmt"

func main() {
	// Создаем карту с элементами
	myMap := map[string]int{
		"apple":  5,
		"banana": 10,
		"orange": 3,
	}

	fmt.Println("Before deletion:", myMap)

	// Удаляем элемент с ключом "banana"
	delete(myMap, "banana")

	fmt.Println("After deletion:", myMap)

	// Проверяем наличие ключа "banana"
	value, exists := myMap["banana"]
	if exists {
		fmt.Println("Value of 'banana' key:", value)
	} else {
		fmt.Println("'banana' key does not exist")
	}

	// Проверяем наличие ключа "apple"
	value, exists = myMap["apple"]
	if exists {
		fmt.Println("Value of 'apple' key:", value)
	} else {
		fmt.Println("'banana' key does not exist")
	}

	// any map
	anyMap := make(map[interface{}]interface{})
	anyMap[1] = "one"
	anyMap["two"] = 2
	for key, value := range anyMap {
		fmt.Println("key: ", key)
		fmt.Println("value: ", value)
	}
}
