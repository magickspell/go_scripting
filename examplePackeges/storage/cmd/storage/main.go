package main

import (
	"fmt"
	"log"

	"gostorage/internal/storage"
)

func main() {
	st := storage.NewStorage()

	fmt.Println("wokr")
	fmt.Println(st)

	file, err := st.Upload("test.txt", []byte("hello"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("uploaded", file)

	foundFile, err := st.GetByID(file.ID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("it is restored", foundFile)
}
