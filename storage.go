package main

import (
	"encoding/json"
	"os"
	"path/filepath"
)

var tasks []Task

func Loadtasks() ([]Task, error) {
	path, err := TasksFilePath()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		return nil, err
	}
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func SaveTasks(task []Task) error {
	data, err := json.MarshalIndent(task, "", "  ")
	if err != nil {
		return err
	}

	path, err := TasksFilePath()
	if err != nil {
		return err
	}

	err = os.WriteFile(path, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func TasksFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	dataDir := filepath.Join(home, ".local", "share", "timo")

	err = os.MkdirAll(dataDir, 0755)
	if err != nil {
		return "", err
	}

	return filepath.Join(dataDir, "tasks.json"), nil
}
