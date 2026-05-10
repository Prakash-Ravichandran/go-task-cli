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

	var allTasks []Task

	existingData, err := os.ReadFile(filename)

	if err == nil && len(existingData) > 0 {
		err := json.Unmarshal(existingData, &allTasks)
		if err != nil {
			fmt.Println("Error reading existing JSON:", err)
		}
	}

	var newTask Task
	task_cli := flag.String("task-cli", "add", "a task for the user can add/update/delete task")
	task := flag.String("add", "todo", "a task for the user can add to a task")
	task_id := flag.Int("id", 0, "an identifer for the task") // TODO: remove it for add
	// task_description := flag.String("")
	flag.Parse()
	fmt.Printf("given command is %s \n", *task_cli)
	fmt.Printf("given task is %s \n", *task)
	fmt.Printf("id is %d \n", *task_id)

	newTaskID := len(allTasks) + 1

	fmt.Println("newTaskID", newTaskID)

	newTask = Task{ID: &newTaskID, Description: task, Status: "", CreatedAt: "", UpdatedAt: ""}

	allTasks = append(allTasks, newTask)

	data, _ := json.MarshalIndent(allTasks, "", "  ") // convert []byte

	os.WriteFile(filename, data, 0644)
	defer file.Close()
}

// 1. get type of task - add, update, delete, mark-in-progress, mark-done, list
// go run main.go --task-cli=add -add "sell mango" -id 3
