package main

import "fmt"

type Task struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"createdat"`
	UpdatedAt   string `json:"updatedat"`
}

func main() {
	fmt.Println("Hello, World!")
}

// Need a JSON to track a task
