package main

import (
	"fmt"
)

func main() {
	mapNumStr := make(map[int64]string)
	fmt.Println(mapNumStr)
	mapNumStr[1] = "one"
	mapNumStr[2] = "two"
	fmt.Println(mapNumStr)
	mapNumStr[1] = "three"
	fmt.Println(mapNumStr)
	mapNumStr[1] = "one"
	mapNumStr[3] = "three"
	mapNumStr[4] = "four"

	// !!! порядок обработки не гарантируется, при разных запусках порядок вывода может менятся
	for key, value := range mapNumStr {
		// fmt.Println(key, value)
		// fmt.Printf(`mapNumStr["%s"] = %d`+"\n", key, value)
		fmt.Printf("mapNumStr[%d] = %s\n", key, value)
	}
}
