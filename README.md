# Task Tracker

Task-tracker is a CLI tool designed programmed in Go to help you track and manage your tasks efficiently. This project is part of a series of tutorials from [roadmap.sh](https://roadmap.sh/projects/task-tracker) and demonstrates how to create a simple task manager with various functionalities.

## Features

- Create, update, delete, and list tasks
- Mark tasks as "done" or "in progress"
- Filter Lists by "To do", "Done" or "In Progress"

## Usage

```
task-tracker [command]
```

### Available Commands
```
-   create: Create a new task
-   delete: Delete a task
-   help: Get help about any command
-   list: List all tasks
-   list done: List all tasks with "Done" status
-   list in-progress: List all tasks with "In Progress" status
-   list to-do: List all tasks with "To Do" status
-   mark-done: Mark a task as done
-   mark-in-progress: Mark a task as in progress
-   update: Update a task
```

### Flags

```
-   `-h, --help`: Show help for the task-tracker CLI
```
## Installation

To install Task-tracker, clone this repository and run the following command in the project directory:

```
git clone https://github.com/your-username/task-tracker.git
cd task-tracker
go install
```

## License

This project is licensed under the MIT License - see the LICENSE file for details.
