package main

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct{}

func (c Cat) Speak() string {
	return "Meow!"
}

// вложенные структуры
type Address struct {
	City, Country string
}

type Person struct {
	Name    string
	Age     int
	Address Address // вложенная структура
}

// стандартный метод стурктуры Person для вывода нформаици на экран
func (p Person) String() string {
	return fmt.Sprintf("%s, возраст %d", p.Name, p.Age)
}

func main() {
	animals := []Animal{Dog{}, Cat{}}
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}

	// иницуиализация струтур
	// Прямая инициализация
	p := Person{"Alice", 30, Address{}}
	fmt.Println(p)
	// Инициализация через указатель
	pp := &Person{Name: "Bob", Age: 25, Address: Address{"city", "country"}}
	fmt.Println(pp)
	// Использование new
	pn := new(Person)
	fmt.Println(pn)
	pn.Name = "Charlie"
	pn.Age = 28
	fmt.Println(pn)
}
