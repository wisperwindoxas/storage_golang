package main

import (
	"fmt"
	"log"

	"github.com/wisperwindoxas/storage/internal/storage"
)

func main() {
	storage := storage.NewStorage()
	
	file, err := storage.Upload("test.txt", []byte("test"))
	if err != nil{
		log.Fatal(err)
	}

	fmt.Println("Uploading", file)

}
 