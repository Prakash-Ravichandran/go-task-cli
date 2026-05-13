package main

import (
	"encoding/json"
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

	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		os.WriteFile(filename, []byte("[]"), 0644)
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
	// command: -add
	task := flag.String("add", "", "a task for the user can add to a task")
	// command: -id
	task_id := flag.Int("id", 0, "an identifer for the task") // TODO: remove it for add
	// command: -delete
	task_delete := flag.Int("delete", 0, "delete a task with an id")
	// command: -update
	task_update := flag.Int("update", 0, "update a task with an id and description")
	// command: -mark-in-progress
	task_mark_in_progress := flag.Int("mark-in-progress", 0, "mark a task as in-progress")
	// command: -mark-done
	task_mark_done := flag.Int("mark-done", 0, "mark a task as done")
	// command: -list
	task_list := flag.String("list", "all", "list all tasks which are done")
	flag.Parse()
	fmt.Printf("given command is %s \n", *task_cli)
	fmt.Printf("given task is %s \n", *task)
	fmt.Printf("id is %d \n", *task_id)
	fmt.Printf("update a task is %d \n", *task_update)
	fmt.Printf("delete is %d \n", *task_delete)
	fmt.Printf("mark-in-progress is %d \n", *task_mark_in_progress)
	fmt.Printf("mark-done is %d \n", *task_mark_done)
	fmt.Printf("list all tasks is %s \n", *task_list)

	newTaskID := len(allTasks) + 1
	fmt.Println("newTaskID", newTaskID)

	if *task != "" {
		newTask = Task{ID: newTaskID, Description: *task, Status: "todo", CreatedAt: time.Now().Format(time.RFC3339), UpdatedAt: time.Now().Format(time.RFC3339)}
		allTasks = append(allTasks, newTask)
		SaveToFile(filename, allTasks)
		fmt.Printf("Task added successfully (ID: %d)\n", newTaskID)
		return
	}

	if *task_update != 0 {
		updatedDescription := flag.Arg(0)
		if updatedDescription == "" {
			fmt.Println("Error: Please provide a new description. Example: -update 1 'new text'")
			return
		}
		updatedTasks := updateTaskByID(task_update, updatedDescription, allTasks)
		fmt.Printf("%+v\n", updatedTasks)
		SaveToFile(filename, updatedTasks)
		return
	}

	if *task_delete != 0 {
		updatedTasks := deleteTaskByID(task_delete, allTasks)
		fmt.Printf("%+v\n", updatedTasks)
		SaveToFile(filename, updatedTasks)
		return
	}

	if *task_mark_in_progress != 0 {

		updatedTasks := markTaskInProgressById(task_mark_in_progress, allTasks)
		fmt.Printf("%+v\n", updatedTasks)
		SaveToFile(filename, updatedTasks)
		return
	}
	if *task_mark_done != 0 {

		updatedTasks := markTaskDoneById(task_mark_done, allTasks)
		fmt.Printf("%+v\n", updatedTasks)
		SaveToFile(filename, updatedTasks)
		return
	}

	if *task_list != "" {
		listTasks(allTasks, *task_list)
	}
}

func SaveToFile(filename string, tasks []Task) {
	data, err := json.MarshalIndent(tasks, "", "  ") // convert []byte

	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	os.WriteFile(filename, data, 0644)
}

func listTasks(tasks []Task, filter string) {
	fmt.Println("ID  | Status          | Description")
	fmt.Println("----|-----------------|------------")
	for _, t := range tasks {
		if filter == "" || filter == "all" || t.Status == filter {
			fmt.Printf("%-3d | %-15s | %s\n", t.ID, t.Status, t.Description)
		}
	}
}

func deleteTaskByID(id *int, t []Task) []Task {
	t = slices.DeleteFunc(t, func(task Task) bool {
		return task.ID == *id
	})

	fmt.Printf("deleted %d", *id)
	return t
}
func updateTaskByID(id *int, updatedDescription string, t []Task) []Task {
	for index := range t {
		if t[index].ID == *id {
			t[index].Description = updatedDescription
		}
	}

	fmt.Printf("updated %d", *id)
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
// TODO: handle this - go run main.go --task-cli -add  "buy mobile 4"
