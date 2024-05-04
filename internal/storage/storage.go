package storage

import "fmt"
import "github.com/wisperwindoxas/storage/internal/file"

type Storage struct {}


func NewStorage() *Storage {
	fmt.Println("Some storage")
	return &Storage{}
}


func (s *Storage) Upload(filename string, blob []byte) (*file.File, error) {
	return  file.NewFile(filename, blob)
	// if err != nil {
	// 	return nil, err
	// }


	// return newFile

}