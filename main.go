package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	storage := NewFileStorage("tasks.json")
	handler := NewTaskHandler(storage)

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Error: Description is required for adding a task")
			return
		}
		handler.AddTask(os.Args[2])

	case "update":
		if len(os.Args) < 4 {
			fmt.Println("Error: ID and description are required for updating a task")
			return
		}
		handler.UpdateTask(os.Args[2], os.Args[3])

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Error: ID is required for deleting a task")
			return
		}
		handler.DeleteTask(os.Args[2])

	case "mark-in-progress":
		if len(os.Args) < 3 {
			fmt.Println("Error: ID is required for marking task as in progress")
			return
		}
		handler.MarkInProgress(os.Args[2])

	case "mark-done":
		if len(os.Args) < 3 {
			fmt.Println("Error: ID is required for marking task as done")
			return
		}
		handler.MarkDone(os.Args[2])

	case "list":
		if len(os.Args) > 2 {
			handler.ListTasksByStatus(os.Args[2])
		} else {
			handler.ListAllTasks()
		}

	default:
		fmt.Printf("Error: Unknown command '%s'\n", command)
		printUsage()
	}
}

func printUsage() {
	fmt.Println("Task Tracker CLI - Manage your tasks")
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Println("  task-cli add <description>         Add a new task")
	fmt.Println("  task-cli update <id> <description> Update a task")
	fmt.Println("  task-cli delete <id>               Delete a task")
	fmt.Println("  task-cli mark-in-progress <id>     Mark task as in progress")
	fmt.Println("  task-cli mark-done <id>            Mark task as done")
	fmt.Println("  task-cli list                      List all tasks")
	fmt.Println("  task-cli list <status>             List tasks by status (todo, in-progress, done)")
	fmt.Println("")
	fmt.Println("Examples:")
	fmt.Println("  task-cli add \"Buy groceries\"")
	fmt.Println("  task-cli update 1 \"Buy groceries and cook dinner\"")
	fmt.Println("  task-cli list done")
}
