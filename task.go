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
		Status:      "todo",
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
		fmt.Printf("ID: %d | Description: %s | Status: %s | Created: %s\n", task.ID, task.Description, task.Status, TimeAgo(task.CreatedAt))
	}
}

func DeleteTask(id string) {
	idToDelete, err := strconv.Atoi(id)
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

func UpdateTask(id string, des string) {
	idToFind, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("❌ Error: ", err)
		return
	}
	tasks, err := Loadtasks()

	found := false
	for i := range tasks {
		if tasks[i].ID == idToFind {
			tasks[i].Description = des
			tasks[i].UpdatedAt = time.Now()
			found = true
			break
		}
	}
	if !found {
		fmt.Println("❌ Error: Task Not Found")
		return
	}
	err = SaveTasks(tasks)
	if err != nil {
		fmt.Println("❌ Error: ", err)
		return
	}
	fmt.Println("✅ Task Updated.")
}

func MarkTask(id string, status string) {
	idToFind, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	tasks, err := Loadtasks()

	found := false
	for i := range tasks {
		if tasks[i].ID == idToFind {
			tasks[i].Status = status
			tasks[i].UpdatedAt = time.Now()
			found = true
			break
		}
	}
	if !found {
		fmt.Println("❌ Error: Task Not Found")
		return
	}

	err = SaveTasks(tasks)
	if err != nil {
		fmt.Println("❌ Error: ", err)
		return
	}
}

func TimeAgo(t time.Time) string {
	elapsed := time.Since(t)

	switch {
	case elapsed < time.Minute:
		return "just now"

	case elapsed < time.Hour:
		minutes := int(elapsed.Minutes())
		if minutes == 1 {
			return "1 minute ago"
		}
		return fmt.Sprintf("%d minutes ago", minutes)

	case elapsed < 24*time.Hour:
		hours := int(elapsed.Hours())
		if hours == 1 {
			return "1 hour ago"
		}
		return fmt.Sprintf("%d hours ago", hours)

	default:
		days := int(elapsed.Hours() / 24)
		if days == 1 {
			return "1 day ago"
		}
		return fmt.Sprintf("%d days ago", days)

	}
}
