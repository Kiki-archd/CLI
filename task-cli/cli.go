package main

import (
	"fmt"
	"strconv"
	"time"
)

type TaskHandler struct {
	storage Storage
}

func NewTaskHandler(storage Storage) *TaskHandler {
	return &TaskHandler{storage: storage}
}

func (th *TaskHandler) AddTask(description string) {
	tasks, err := th.storage.LoadTasks()
	if err != nil {
		fmt.Printf("Error loading tasks: %v\n", err)
		return
	}

	id := strconv.Itoa(len(tasks) + 1)
	task := NewTask(id, description)
	tasks = append(tasks, *task)

	if err := th.storage.SaveTasks(tasks); err != nil {
		fmt.Printf("Error saving task: %v\n", err)
		return
	}

	fmt.Printf("Task added successfully (ID: %s)\n", id)
}

func (th *TaskHandler) UpdateTask(id, description string) {
	tasks, err := th.storage.LoadTasks()
	if err != nil {
		fmt.Printf("Error loading tasks: %v\n", err)
		return
	}

	found := false
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Description = description
			tasks[i].UpdatedAt = time.Now()
			found = true
			break
		}
	}

	if !found {
		fmt.Printf("Error: Task with ID %s not found\n", id)
		return
	}

	if err := th.storage.SaveTasks(tasks); err != nil {
		fmt.Printf("Error updating task: %v\n", err)
		return
	}

	fmt.Printf("Task %s updated successfully\n", id)
}

func (th *TaskHandler) DeleteTask(id string) {
	tasks, err := th.storage.LoadTasks()
	if err != nil {
		fmt.Printf("Error loading tasks: %v\n", err)
		return
	}

	found := false
	var updatedTasks []Task
	for _, task := range tasks {
		if task.ID != id {
			updatedTasks = append(updatedTasks, task)
		} else {
			found = true
		}
	}

	if !found {
		fmt.Printf("Error: Task with ID %s not found\n", id)
		return
	}

	if err := th.storage.SaveTasks(updatedTasks); err != nil {
		fmt.Printf("Error deleting task: %v\n", err)
		return
	}

	fmt.Printf("Task %s deleted successfully\n", id)
}

func (th *TaskHandler) MarkInProgress(id string) {
	th.updateStatus(id, StatusInProgress, "in progress")
}

func (th *TaskHandler) MarkDone(id string) {
	th.updateStatus(id, StatusDone, "done")
}

func (th *TaskHandler) updateStatus(id string, status TaskStatus, statusName string) {
	tasks, err := th.storage.LoadTasks()
	if err != nil {
		fmt.Printf("Error loading tasks: %v\n", err)
		return
	}

	found := false
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Status = status
			tasks[i].UpdatedAt = time.Now()
			found = true
			break
		}
	}

	if !found {
		fmt.Printf("Error: Task with ID %s not found\n", id)
		return
	}

	if err := th.storage.SaveTasks(tasks); err != nil {
		fmt.Printf("Error marking task as %s: %v\n", statusName, err)
		return
	}

	fmt.Printf("Task %s marked as %s successfully\n", id, statusName)
}

func (th *TaskHandler) ListAllTasks() {
	tasks, err := th.storage.LoadTasks()
	if err != nil {
		fmt.Printf("Error loading tasks: %v\n", err)
		return
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found")
		return
	}

	fmt.Println("All Tasks:")
	fmt.Println("----------")
	for _, task := range tasks {
		printTask(task)
	}
}

func (th *TaskHandler) ListTasksByStatus(status string) {
	tasks, err := th.storage.LoadTasks()
	if err != nil {
		fmt.Printf("Error loading tasks: %v\n", err)
		return
	}

	var statusFilter TaskStatus
	switch status {
	case "todo":
		statusFilter = StatusTodo
	case "in-progress":
		statusFilter = StatusInProgress
	case "done":
		statusFilter = StatusDone
	default:
		fmt.Printf("Error: Invalid status '%s'. Use: todo, in-progress, or done\n", status)
		return
	}

	var filteredTasks []Task
	for _, task := range tasks {
		if task.Status == statusFilter {
			filteredTasks = append(filteredTasks, task)
		}
	}

	if len(filteredTasks) == 0 {
		fmt.Printf("No tasks with status '%s' found\n", status)
		return
	}

	fmt.Printf("Tasks (%s):\n", status)
	fmt.Println("----------")
	for _, task := range filteredTasks {
		printTask(task)
	}
}

func printTask(task Task) {
	fmt.Printf("ID: %s\n", task.ID)
	fmt.Printf("Description: %s\n", task.Description)
	fmt.Printf("Status: %s\n", task.Status)
	fmt.Printf("Created: %s\n", task.CreatedAt.Format("2006-01-02 15:04:05"))
	fmt.Printf("Updated: %s\n", task.UpdatedAt.Format("2006-01-02 15:04:05"))
	fmt.Println("---")
}
