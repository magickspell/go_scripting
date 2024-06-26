package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type myStruct struct {
	privateField string
}

func (s myStruct) MarshalJSON() ([]byte, error) {
	const prefix, suffix = `{"privateField": "`, `"}`

	buf := new(bytes.Buffer)
	buf.Grow(len(prefix) + len(s.privateField) + len(suffix))

	buf.WriteString(prefix)
	buf.WriteString(s.privateField)
	buf.WriteString(suffix)

	return buf.Bytes(), nil
}

// реализует ли myStruct интерфейс json.Marshaler
var _ json.Marshaler = (*myStruct)(nil)

func main() {
	fmt.Println()

	var data = myStruct{
		privateField: "hello world 111",
	}

	jsonBytes, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonBytes))
}
