package main

import "fmt"

type PolInterface interface {
	IncMethod(in string) (out string)
}

type PolStruct struct {
	prefix string
}

func (s incStruct) PolMethod(in string) (out string) {
	return s.prefix + " : " + in
}

func PolSomeInterface(prefix string) IncInterface {
	return incStruct{prefix}
}

func main() {
	v := PolSomeInterface("one")
	fmt.Println(v.IncMethod("two")) // Output: onetwo
}
