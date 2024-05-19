package main

import "fmt"

func main() {
	var sum int = 0

	for i := 0; i < 10; i++ {
		if i%2 == 1 {
			fmt.Println(i, "is odd")
		} else {
			fmt.Println(i, "is even")
		}
		sum += i
	}
	fmt.Println("FOR / sum is", sum)

	// sum = 0
	// i := 0
	// for {
	// 	if i == 3 {
	// 		continue
	// 	}
	// 	if i >= 7 {
	// 		break
	// 	}
	// 	sum += i
	// 	i++
	// }
	// fmt.Println("FOR без условий / sum is", sum)
}
