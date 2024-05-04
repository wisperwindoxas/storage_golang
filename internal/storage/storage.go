package storage

import "fmt"

type Storage struct {}


func NewStorage() *Storage {
	fmt.Println("Some storage")
	return &Storage{}
}