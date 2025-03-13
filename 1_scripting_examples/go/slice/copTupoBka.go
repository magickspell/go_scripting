package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	people := []Person{
		{"Bob", 31},
		{"Alice", 52},
		{"Charlie", 23},
		{"Gerry", 77},
		{"Prestone", 15},
	}

	fmt.Println(people)

	// Сортировка по полю Name
	sort.Slice(people, func(i, j int) bool {
		return people[i].Name < people[j].Name
	})

	fmt.Println(people) // Выведет отсортированный по имени список: [Alice Bob Charlie]

	// Сортировка по полю Age
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println(people)
}
