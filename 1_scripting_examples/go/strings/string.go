package main

import (
	"fmt"
	"strings"
)

func add(x int, y int) int {
	return x + y
}

func str(a, b string) (string, string) {
	return a + b, b + a
}

// naked return
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// costa := false // := нельзя юзать в глобальном окружении
var constanta bool = false

func main() {
	fmt.Println(add(42, 13))
	fmt.Println(str("a", "b"))
	fmt.Println(split(25))

	// nil value
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)

	// pointers
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j

	fmt.Println("ARRAY and SLICE")
	var array = [3]int{20, 1, 99}
	fmt.Println(array[0])
	fmt.Println(array)
	var slice = []string{"hello"}
	fmt.Println(slice)
	slice = append(slice, "world")
	fmt.Println(slice)
	slice = (&[4]string{"1", "2", "3", "4"})[1:3]
	fmt.Println(slice)
	slice = slice[:]
	fmt.Println(slice)
	fmt.Println("capacity: ", cap(slice))
	fmt.Println("length: ", len(slice))
	// nil value of a maded slice will be 0 for ints
	sliceMaded := make([]int, 4)
	fmt.Println(sliceMaded)
	for index, value := range sliceMaded {
		fmt.Printf("index: %v, value: %v \n", index, value)
	}

	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	fmt.Println("MAPS")
	mapp := make(map[string]int)
	fmt.Println(mapp)
	mapp["two"] = 2
	fmt.Println(mapp)
	delete(mapp, "two")
	fmt.Println(mapp)
	element, ok := mapp["two"]
	fmt.Println("The value:", element, "Present?", ok)

	// fibanachi
	fibonacci := func() func() int {
		a, b := 0, 1
		return func() int {
			result := a
			a, b = b, a+b
			return result
		}
	}
	for i := 0; i < 10; i++ {
		fmt.Println(fibonacci()())
	}
}
