package main

import "fmt"

type IncInterface interface {
	IncMethod(in string) (out string)
}

type incStruct struct { // приватная
	prefix string
}

func (s incStruct) IncMethod(in string) (out string) {
	return s.prefix + in
}

func NewSomeInterface(prefix string) IncInterface {
	return incStruct{prefix}
}

func main() {
	v := NewSomeInterface("one")
	fmt.Println(v.IncMethod("+two")) // Output: onetwo
}
