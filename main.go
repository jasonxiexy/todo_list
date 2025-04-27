package main

import (
	"fmt"
)

func main() {
	list, err := loadData()
	if err != nil {
		fmt.Printf("❌ %s\n", err)
		return
	}

	switch {
	case islistTask:
		printTask(false, list)
	case islistAllTask:
		printTask(true, list)
	case doneTaskID != 0:
		doneTask(list)
	case addTaskTitle != "":
		addTask(list)
	case deleteTaskID != 0:
		deleteTask(list)
	}
}
