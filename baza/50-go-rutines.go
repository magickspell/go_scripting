package main

import (
	"fmt"
	"runtime"
	"time"
)

// func main() { // выведет в разнобой
// 	for i := 0; i < 5; i++ {
// 		fmt.Println("outer i: ", i)
// 		go func() {
// 			time.Sleep(time.Second)
// 			fmt.Println(i)
// 		}()
// 	}
// 	time.Sleep(2 * time.Second)
// }

func main() { // тоже выведет в разнобой, но мы уже четко передали переменную, нужно делать так!
	// runtime.GOMAXPROCS(1) // оставляет только 1 рабочий процессор для выполнения кода

	for i := 0; i < 10; i++ {
		// fmt.Println("outer i: ", i)
		go func() {
			// time.Sleep(time.Second)
			// fmt.Println(counter)
			fmt.Println("\n gorutine %v", i)
		}()
		runtime.Gosched() // передает управление другой горутине, упорядочивает вызовы
	}
	time.Sleep(1 * time.Second)
}
