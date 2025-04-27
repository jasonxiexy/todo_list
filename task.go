package main

import (
	"flag"
	"fmt"
	"time"
)

// Command-line flags
var addTaskTitle string
var doneTaskID int
var deleteTaskID int
var islistTask bool
var islistAllTask bool

func init() {
	flag.StringVar(&addTaskTitle, "add", "", "Add New Task")
	flag.IntVar(&doneTaskID, "done", 0, "Finish Tasks")
	flag.IntVar(&deleteTaskID, "del", 0, "Delete Tasks")
	flag.BoolVar(&islistTask, "list", false, "List unfinished tasks")
	flag.BoolVar(&islistAllTask, "all", false, "List all tasks")
	flag.Parse()
}

// Structs
type todoTask struct {
	ID        int       `json:"id"`
	TaskTitle string    `json:"task_title"`
	CreatedAt time.Time `json:"created_at"`
	IsDone    bool      `json:"is_done"`
	IsDeleted bool      `json:"is_deleted"`
}

type todoList []todoTask

// Add Task
func addTask(list todoList) {
	task := todoTask{
		ID:        len(list) + 1,
		TaskTitle: addTaskTitle,
		CreatedAt: time.Now(),
		IsDone:    false,
	}
	list = append(list, task)
	err := saveData(list)
	if err != nil {
		fmt.Printf("❌ %s\n", err)
	} else {
		fmt.Println("✅ Successfully added task")
	}
}

// Complete Task
func doneTask(list todoList) {
	task := &list[doneTaskID-1]
	task.IsDone = true
	err := saveData(list)
	if err != nil {
		fmt.Printf("❌ %s\n", err)
	} else {
		fmt.Println("✅ Successfully finished task")
	}
}

// Delete Task
func deleteTask(list todoList) {
	task := &list[deleteTaskID-1]
	task.IsDeleted = true
	err := saveData(list)
	if err != nil {
		fmt.Printf("❌ %s\n", err)
	} else {
		fmt.Println("✅ Successfully remove task")
	}
}
