package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	fmt.Println()
	{ // MARSHAL - парсит в строчку
		data := map[string]interface{}{
			"a": 123,
			"b": "text",
			"c": struct{}{},
			"d": 123.456,
		}

		jsonBytes, err := json.Marshal(data)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(jsonBytes))
	}

	fmt.Println()
	{ // MARSHAL INDENT - парсит и вставляет префиксы и инденты
		data := map[string]interface{}{
			"a": 123,
			"b": "text",
			"c": struct{}{},
			"d": 123.456,
			"e": 2,
		}

		jsonBytes, err := json.MarshalIndent(data, "<prefix>", "<indent>")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(jsonBytes))
	}
}
