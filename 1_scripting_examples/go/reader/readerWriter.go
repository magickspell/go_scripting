package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	reader := strings.NewReader("Hello, Reader!")

	bytes := make([]byte, 8)
	for {
		valueByteLength, err := reader.Read(bytes)
		fmt.Printf("valueByteLength = %v err = %v bytes = %v\n", valueByteLength, err, bytes)
		fmt.Printf("bytes[:valueByteLength] = %q\n", bytes[:valueByteLength])
		if err == io.EOF {
			break
		}
	}
}
