package main

import (
	"fmt"
	"strconv"
	"time"
)

type Task struct {
	ID          int
	Description string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func AddTask(description string) {
	tasks, err := Loadtasks()
	if err != nil {
		fmt.Println(err)
		return
	}
	maxID := 0
	for _, task := range tasks {
		if task.ID > maxID {
			maxID = task.ID
		}
	}
	newTask := Task{
		ID:          maxID + 1,
		Description: description,
		Status:      "Todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	tasks = append(tasks, newTask)
	err = SaveTasks(tasks)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("✅ Task Added.")
}

func ListTasks() {
	tasks, err := Loadtasks()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, task := range tasks {
		fmt.Printf("ID: %d | Description: %s | Status: %s\n", task.ID, task.Description, task.Status)
	}
}

func DeleteTask(id string) {
	idStr := id
	idToDelete, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("❌ ID Should a Int: ", err)
		return
	}
	tasks, err := Loadtasks()
	if err != nil {
		fmt.Println("❌ Error: ", err)
		return
	}
	newTasks := make([]Task, 0, len(tasks))
	found := false
	for _, task := range tasks {
		if task.ID == idToDelete {
			found = true
			continue
		}
		newTasks = append(newTasks, task)
	}
	if !found {
		fmt.Println("❌ ID Not Found.")
		return
	}
	err = SaveTasks(newTasks)
	if err != nil {
		fmt.Println("❌ Error: ", err)
		return
	}
	fmt.Println("✅ Task Deleted.")
}
