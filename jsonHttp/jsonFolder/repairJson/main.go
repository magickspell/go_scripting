package main

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

// Contract
type Contract struct {
	Posts []post `json:"posts"`
}

type post struct {
	ID    id     `json:"id"`
	Title string `json:"title"`
}

type id int64

func (i *id) UnmarshalJSON(b []byte) error {
	if b[0] == '"' {
		b = bytes.Trim(b, `"`)
	}

	v, err := strconv.ParseInt(string(b), 10, 64)
	if err != nil {
		return err
	}

	*i = id(v)
	return nil
}

//go:embed data.json
var jsonBytes []byte

func main() {
	fmt.Println()

	var data = new(Contract)
	err := json.Unmarshal(jsonBytes, data)
	if err != nil {
		log.Fatal(err)
	}

	for _, p := range data.Posts {
		fmt.Printf("%d: %s\n", p.ID, p.Title)
	}
}
