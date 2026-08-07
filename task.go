package main

import (
	"fmt"
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
	newTask := Task{
		ID:          len(tasks) + 1,
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

}

func ListTasks() {
	tasks, err := Loadtasks()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, task := range tasks {
		fmt.Printf("ID: %d | %s | %s\n", task.ID, task.Description, task.Status)
	}
}
