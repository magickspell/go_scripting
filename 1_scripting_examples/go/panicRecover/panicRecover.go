package main

import "fmt"

func test() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic message 2: ", err)
			fmt.Println("panic recovered twice (defer всплывают)")
		}
	}()
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic message 1: ", err)
			fmt.Println("panic recovered")
			panic("panic2")
		}
	}()
	fmt.Println("somethinf happening")
	panic("ooops, PANIC!")
	fmt.Println("somethinf happening after panic") // не выполняется
}

func main() {
	test()
}
