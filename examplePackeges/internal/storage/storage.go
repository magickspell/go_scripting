package storage

import "gostorage/internal/file"

type Storage struct{}

func NewStorage() *Storage { // можно без показателей, но лучше вот так
	return &Storage{}
}

func (s *Storage) Upload(filename string, blob []byte) (*file.File, error) {
	return file.NewFile(filename, blob)
}
