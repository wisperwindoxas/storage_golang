package storage

import "fmt"
import "github.com/wisperwindoxas/storage/internal/file"

type Storage struct {
	File map[uuid.UUID] *file.File
}


func NewStorage() *Storage {
	return &Storage{
		
	}
}


func (s *Storage) Upload(filename string, blob []byte) (*file.File, error) {
	return  file.NewFile(filename, blob)
	// if err != nil {
	// 	return nil, err
	// }


	// return newFile

}

func (s *Storage) GetByID(fileID uuuid.UUID) (*file.File, error) {
	
}