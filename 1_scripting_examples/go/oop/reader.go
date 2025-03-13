package main

import "io"

type Reader interface {
	Read(p []byte) (n int, err error)
}

func readData(r io.Reader) ([]byte, error) {
	buf := make([]byte, 1024) // Создаем буфер размером 1024 байта
	n, err := r.Read(buf)
	if err != nil {
		return nil, err
	}
	return buf[:n], nil
}

func main() {
	/**/
}
