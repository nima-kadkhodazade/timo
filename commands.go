package main

import (
	"fmt"
	"os"
)

func ManageCommands() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: timo <command>")
		return
	}
	command := os.Args[1]
	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("❌ Error: Please provide a task description")
			return
		}
		AddTask(os.Args[2])
	case "list":
		ListTasks()
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("❌ Error: Please Enter Your ID Task")
			return
		}
		DeleteTask(os.Args[2])
	case "update":
		if len(os.Args) < 3 {
			fmt.Println("❌ Error: Please Enter Your ID Task")
			return
		}
		if len(os.Args) < 4 {
			fmt.Println("❌ Error: Please Enter Your Description")
			return
		}
		UpdateTask(os.Args[2], os.Args[3])

	default:
		fmt.Println("❌ Wrong Command...!")
	}
}
