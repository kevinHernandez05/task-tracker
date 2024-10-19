package utils

import (
	"encoding/json"
	"io"
	"os"
	"task-tracker/entity"
	"time"
)

func LoadTasks(filter string) ([]entity.Task, error) {
	file, err := os.Open("tasks.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var tasks []entity.Task

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tasks)

	if err != nil {
		if err == io.EOF {
			return []entity.Task{}, nil
		}
		return nil, err
	}

	if filter != "" {
		var filteredTasks []entity.Task
		for _, task := range tasks {
			if task.Status == filter {
				filteredTasks = append(filteredTasks, task)
			}
		}
		tasks = filteredTasks
	}

	return tasks, nil
}

func SaveTasks(tasks []entity.Task) error {
	file, err := os.Create("tasks.json")
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(tasks)
}

func DeleteTaskById(Task []entity.Task, id int) ([]entity.Task, int) {

	index := -1 //-1

	for i, task := range Task {
		if task.Id == id {
			index = i
			break
		}
	}

	if index == -1 {
		return []entity.Task{}, index
	}

	Task = append(Task[:index], Task[index+1:]...)

	return Task, index
}

func UpdateTaskById(Description string, id int, Task []entity.Task) ([]entity.Task, int) {

	index := -1

	for i, task := range Task {
		if task.Id == id {
			Task[i].Description = Description
			Task[i].UpdatedAt = time.Now().Format(time.RFC822)

			index = i
			break
		}
	}

	return Task, index
}

func UpdateTaskStatusById(Status string, id int, Task []entity.Task) ([]entity.Task, int) {

	index := -1

	for i, task := range Task {
		if task.Id == id {
			Task[i].Status = Status
			Task[i].UpdatedAt = time.Now().Format(time.RFC822)

			index = i
			break
		}
	}

	return Task, index
}

func GetNextId(tasks []entity.Task) int {
	maxId := 0
	for _, task := range tasks {
		if task.Id > maxId {
			maxId = task.Id
		}
	}
	return maxId + 1
}
