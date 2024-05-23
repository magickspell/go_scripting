package storage

import (
	"fmt"
	"gostorage/internal/file"

	"github.com/google/uuid"
)

type Storage struct {
	files map[uuid.UUID]*file.File // приватная переменная - инкапсулированная в структуре
}

func NewStorage() *Storage { // можно без показателей, но лучше вот так
	return &Storage{
		files: make(map[uuid.UUID]*file.File),
	}
}

func (s *Storage) Upload(filename string, blob []byte) (*file.File, error) {
	newFile, err := file.NewFile(filename, blob)
	if err != nil {
		return nil, err
	}

	s.files[newFile.ID] = newFile

	return newFile, nil

}

func (s *Storage) GetByID(fileID uuid.UUID) (*file.File, error) {
	foundFile, ok := s.files[fileID]
	if !ok {
		// return nil, errors.New(fmt.Sprintf("file %v not found", fileID))
		return nil, fmt.Errorf("file %v not found", fileID)
	}

	return foundFile, nil
}
