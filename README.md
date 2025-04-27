# 📝 Todo List CLI (Golang)

A simple command-line Todo List application written in Go, using a local JSON file for task storage.

## Features

- Add new tasks
- Complete tasks
- Delete tasks
- View unfinished tasks
- View all tasks (including completed and deleted)

## Project Structure
```
todo_List/
│
├── go.mod        # dependency
├── main.go       # entry point
├── task.go       # defines task struct and task operations
├── storage.go    # load/save tasks from/to a JSON file
├── print.go      # functions to print the tasks
└── tasks.json    # (auto-created when you run your program)
```


## 🛠 Setup Instructions

Follow these steps to set up and run the project locally:

1. **Get Started**

    Clone project:

    ```bash
    git clone 
    cd todo_List
    ```

1. **Tidy Up Modules**

    Clean and verify dependencies:

    ```bash
    go mod tidy
    ```

2. **Build the Project**

    Compile the program:

    ```bash
    go build -o todoList.exe
    ```

3. **Run Commands**

    ```bash
    ./todoList.exe -h                   # list all available commands
    ```
    Examples:

    ```bash
    ./todoList.exe -add "Run for 10km"   # Add a new task
    ./todoList.exe -list              # View unfinished tasks
    ./todoList.exe -done 1             # Mark task 1 as done
    ./todoList.exe -del 2              # Delete task 2
    ./todoList.exe -all                # View all tasks
    ```

---
