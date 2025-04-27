package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

const TASKS_FILE_NAME = "tasks.json"

// Load task data
func loadData() (todoList, error) {
	var list todoList
	dataFile, err := os.OpenFile(TASKS_FILE_NAME, os.O_CREATE|os.O_RDONLY, os.ModePerm)
	if err != nil {
		return list, fmt.Errorf("Failed to open file: %s", err)
	}
	defer dataFile.Close()

	data, err := io.ReadAll(dataFile)
	if err != nil {
		return list, fmt.Errorf("Failed to read file: %s", err)
	}

	if len(data) != 0 {
		err = json.Unmarshal(data, &list)
		if err != nil {
			return list, fmt.Errorf("Failed to read json data: %s", err)
		}
	}

	return list, nil
}

// Save task data
func saveData(list todoList) error {
	dataFile, err := os.OpenFile(TASKS_FILE_NAME, os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return fmt.Errorf("Failed to open file: %s", err)
	}
	defer dataFile.Close()

	data, err := json.Marshal(list)
	if err != nil {
		return fmt.Errorf("Failed to read json data: %s", err)
	}

	_, err = dataFile.Write(data)
	if err != nil {
		return fmt.Errorf("Failed to write data: %s", err)
	}

	return nil
}
