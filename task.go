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
	task := Task{
		ID:          1,
		Description: description,
		Status:      "Todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	fmt.Printf("%+v\n", task)
}
