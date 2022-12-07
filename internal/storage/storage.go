package storage

import (
	"fmt"

	"github.com/DeadLandscape/course_storage/internal/file"

	"github.com/google/uuid"
)

type Storage struct {
	files map[uuid.UUID]*file.File
}

func NewStorage() *Storage {
	return &Storage{
		files: make(map[uuid.UUID]*file.File),
	}
}

func (s *Storage) Upload(filename string, blob []byte) (*file.File, error) {
	file, err := file.NewFile(filename, blob)
	if err != nil {
		return nil, err
	}

	s.files[file.ID] = file

	return file, nil

}

func (s *Storage) GetFile(id uuid.UUID) (*file.File, error) {
	target, ok := s.files[id]
	if !ok {
		//return nil, errors.New(fmt.Sprintf("Fine %v not found", id))
		return nil, fmt.Errorf("file %v not found", id)
	}
	return target, nil
}
