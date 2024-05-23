package file

import (
	"fmt"

	"github.com/google/uuid"
)

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
		Name: filename, // все переменные публичные
		Data: blob,
		ID:   fileID,
	}, nil
}

// создание строкового предстваления структуры
func (f *File) String() string { // если такой метод есть по дефолту - другие модули будут использовать его для красивого вывода
	return fmt.Sprintf("File(%s, %v)", f.Name, f.ID)
}
