package main

import (
	"errors"
	"fmt"
	"os"
)

type Task struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"createdat"`
	UpdatedAt   string `json:"updatedat"`
}

func main() {
	filename := "tasks.json"

	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0644)

	if err != nil {
		if errors.Is(err, os.ErrExist) {
			fmt.Println("File already exists")
		} else {
			fmt.Printf("Err in opening the file %v\n", err)
		}
	}
	defer file.Close()
}
