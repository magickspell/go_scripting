package file

import "github.com/google/uuid"

type File struct {
	ID   uuid.UUID
	Name string
	Data []byte
}

func NewFile(filename string, blob []byte) (*File, error) {
	fileID, err := uuid.NewUUID()
	if err != nil { // обычно в go так работают с ошибками, т.е. возвращают ошибку из функции, после всех нужных объектов
		return nil, err
	}

	return &File{
		Name: filename,
		Data: blob,
		ID:   fileID,
	}, nil
}
