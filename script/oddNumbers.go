package main

import (
	"encoding/json"
	"fmt"
)

type Answer struct {
	Amount int
	Status string
}

func OddCount(n int) int {
	var counter int = 0

	for i := 1; i <= n; i++ {
		if i%2 != 0 {
			counter++
		}
	}

	return counter
}

func main() {
	fmt.Println("start")
	var result int = OddCount(25)
	fmt.Println(result)
	var answer Answer = Answer{
		Amount: result,
		Status: "success",
	}
	fmt.Println(answer)
	var jsonAnswer, err = json.Marshal(answer)
	fmt.Println(err)
	fmt.Println(string(jsonAnswer))
	fmt.Println("end")
}
