package main

import (
	"fmt"

	"github.com/wisperwindoxas/storage/internal/storage"
)

func main() {
	storage := storage.NewStorage()
	fmt.Println(storage)
	fmt.Println("Starting")
}
 