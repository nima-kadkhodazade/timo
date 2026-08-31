package main

import (
	"fmt"
	"os"
	"strings"
)

func ManageCommands() {
	if len(os.Args) < 2 {
		fmt.Println("⚠️  Usage: timo <command>")
		return
	}
	if strings.HasPrefix(os.Args[1], "mark-") {
		status := strings.TrimPrefix(os.Args[1], "mark-")
		if status != "todo" && status != "in-progress" && status != "done" {
			fmt.Println("❌ Error: Invalid status")
			return
		}

		if len(os.Args) < 3 {
			fmt.Println("❌ Error: Please provide a task ID")
			return
		}
		id := os.Args[2]
		MarkTask(id, status)
		return
	}
	command := os.Args[1]
	switch command {
	case "--help":
		ShowHelp()
	case "-v", "--version":
		ShowVersion()
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
		fmt.Println("timo: try 'timo --help' for more information")
	}
}

func ShowHelp() {
	fmt.Println(`	
Usage:
  timo <command> [arguments]
	
Commands:
  add <description>             Add a new task
  list                          List all tasks
  update <id> <description>     Update a task description
  delete <id>                   Delete a task
  mark-todo <id>                Mark task as todo
  mark-in-progress <id>         Mark task as in-progress
  mark-done <id>                Mark task as done
  --help                        Show this help message
  -v, --version                 Show Timo version
	
Examples:
  timo add "Learn Go"
  timo list
  timo update 1 "Learn Go and build Timo"
  timo mark-in-progress 1
  timo mark-done 1
  timo delete 1`)
}

func ShowVersion() {
	fmt.Println(version)
}
