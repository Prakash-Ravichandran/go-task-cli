package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
)

type Task struct {
	ID          *int    `json:"id"`
	Description *string `json:"description"`
	Status      string  `json:"status"`
	CreatedAt   string  `json:"createdat"`
	UpdatedAt   string  `json:"updatedat"`
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

	var firstItem Task
	task_cli := flag.String("task-cli", "add", "a task for the user can add/update/delete task")
	task := flag.String("add", "todo", "a task for the user can add to a task")
	task_id := flag.Int("id", 0, "an identifer for the task")
	// task_description := flag.String("")
	flag.Parse()
	fmt.Printf("given command is %s \n", *task_cli)
	fmt.Printf("given task is %s \n", *task)
	fmt.Printf("id is %d \n", *task_id)

	firstItem = Task{ID: task_id, Description: task, Status: "", CreatedAt: "", UpdatedAt: ""}

	data, _ := json.MarshalIndent(firstItem, "", "  ")

	os.WriteFile(filename, data, 0644)
	defer file.Close()
}

// 1. get type of task - add, update, delete, mark-in-progress, mark-done, list
