package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"slices"
	"time"
)

type Task struct {
	ID          int    `json:"id"`
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
	// command: delete
	task_delete := flag.Int("delete", 0, "delete a task with an id")
	// command: mark-in-progress
	task_mark_in_progress := flag.Int("mark-in-progress", 0, "mark a task as in-progress")
	// command: mark-done
	task_mark_done := flag.Int("mark-done", 0, "mark a task as done")
	// command: list todo
	task_list := flag.String("list", "all", "list all tasks which are done")
	flag.Parse()
	fmt.Printf("given command is %s \n", *task_cli)
	fmt.Printf("given task is %s \n", *task)
	fmt.Printf("id is %d \n", *task_id)
	fmt.Printf("delete is %d \n", *task_delete)
	fmt.Printf("mark-in-progress is %d \n", *task_mark_in_progress)
	fmt.Printf("mark-done is %d \n", *task_mark_done)
	fmt.Printf("list all tasks is %s \n", *task_list)

	newTaskID := len(allTasks) + 1

	fmt.Println("newTaskID", newTaskID)

	newTask = Task{ID: newTaskID, Description: *task, Status: "todo", CreatedAt: time.Now().Format(time.RFC3339), UpdatedAt: time.Now().Format(time.RFC3339)}

	allTasks = append(allTasks, newTask)

	if *task_delete != 0 {
		updatedTasks := deleteTaskByID(task_delete, allTasks)
		fmt.Printf("%+v\n", updatedTasks)
		data, _ := json.MarshalIndent(updatedTasks, "", "  ") // convert []byte
		os.WriteFile(filename, data, 0644)
		os.Exit(1)
	}

	if *task_mark_in_progress != 0 {

		updatedTasks := markTaskInProgressById(task_mark_in_progress, allTasks)
		fmt.Printf("%+v\n", updatedTasks)
		data, _ := json.MarshalIndent(updatedTasks, "", "  ") // convert []byte
		os.WriteFile(filename, data, 0644)
		os.Exit(1)
	}
	if *task_mark_done != 0 {

		updatedTasks := markTaskDoneById(task_mark_done, allTasks)
		fmt.Printf("%+v\n", updatedTasks)
		data, _ := json.MarshalIndent(updatedTasks, "", "  ") // convert []byte
		os.WriteFile(filename, data, 0644)
		os.Exit(1)
	}

	if *task_list != "" {

		if *task_list == "all" {
			var tasktilte []string
			for index := range allTasks {
				tasktilte = append(tasktilte, allTasks[index].Description)
			}

			fmt.Println("all tasks list", allTasks)
			os.Exit(1)
		}
	}

	if *task_list != "" {

		if *task_list == "done" {
			var taskDone []Task
			for index := range allTasks {
				if allTasks[index].Status == "done" {
					taskDone = append(taskDone, allTasks[index])
				}
			}

			fmt.Println("all done list", taskDone)
			os.Exit(1)
		}
	}

	if *task_list != "" {

		if *task_list == "in-progress" {
			var tasksInProgress []Task
			for index := range allTasks {
				if allTasks[index].Status == "in-progress" {
					tasksInProgress = append(tasksInProgress, allTasks[index])
				}
			}

			fmt.Println("in progress list", tasksInProgress)
			os.Exit(1)
		}
	}

	data, _ := json.MarshalIndent(allTasks, "", "  ") // convert []byte

	os.WriteFile(filename, data, 0644)

	fmt.Printf("Task added successfully (ID: %d)", newTaskID)
	defer file.Close()
}

func deleteTaskByID(id *int, t []Task) []Task {
	t = slices.DeleteFunc(t, func(task Task) bool {
		return task.ID == *id
	})

	fmt.Printf("deleted %d", *id)
	return t
}

func markTaskInProgressById(id *int, tasks []Task) []Task {

	for index := range tasks {
		if tasks[index].ID == *id {
			tasks[index].Status = "in-progress"
			tasks[index].UpdatedAt = time.Now().String()
		}
	}
	fmt.Printf("updated %d", *id)
	return tasks
}

func markTaskDoneById(id *int, tasks []Task) []Task {

	for index := range tasks {
		if tasks[index].ID == *id {
			tasks[index].Status = "done"
			tasks[index].UpdatedAt = time.Now().String()
		}
	}
	fmt.Printf("updated as done %d", *id)
	return tasks
}

// 1. get type of task - add, update, delete, mark-in-progress, mark-done, list
// go run main.go --task-cli=add -add "sell mango" -id 3

// TODO: reove id in flag
// TODO: handle same task already exists
// TODO: prints the address here fmt.Printf("%+v\n", updatedTasks)
