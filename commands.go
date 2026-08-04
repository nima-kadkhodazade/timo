package main

import (
	"fmt"
	"os"
)

func ManageCommands() {
	if len(os.Args) < 2 {
		fmt.Println("Error")
		fmt.Println("Usage: task-tracker <command>")
		return
	}
	command := os.Args[1]
	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Error: Please provide a task description")
			return
		}
		AddTask(os.Args[2])
	case "list":
		fmt.Println("List All Tasks...")
	case "delete":
		fmt.Println("Task Deleted...")
	case "update":
		fmt.Println("Task Updated...")

	default:
		fmt.Println("Wrong Command...!")
	}
}
