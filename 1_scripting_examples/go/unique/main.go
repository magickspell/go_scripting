package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// запуск тестов -go test -v ./..
// cat sort.txt | go run main.go - запуск ПО

func unique(input io.Reader, output io.Writer) error {
	in := bufio.NewScanner(input)
	var prev string
	for in.Scan() {
		txt := in.Text()
		if txt == prev {
			continue
		}
		if txt < prev {
			return fmt.Errorf("file not sorted")
		}
		prev = txt
		fmt.Fprintln(output, txt)
	}
	return nil
}

func main() {
	err := unique(os.Stdin, os.Stdout)
	if err != nil {
		panic(err.Error())
	}
}
