/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"task-tracker/entity"
	"task-tracker/utils"
	"time"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new task",
	Long:  `Create a new task that you want to track.`,
	Run: func(cmd *cobra.Command, args []string) {
		createATask(args[0])
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}

func createATask(Description string) {

	// Cargar las tareas actuales
	tasks, err := utils.LoadTasks("")
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	//create the new task
	task := entity.Task{
		Id:          utils.GetNextId(tasks),
		Description: Description,
		Status:      "To do",
		CreatedAt:   time.Now().Format(time.RFC822),
		UpdatedAt:   "",
	}

	tasks = append(tasks, task)

	err = utils.SaveTasks(tasks)

	if err != nil {
		fmt.Println("Error saving tasks:", err)
		return
	}

	fmt.Println("Task created successfully")

}
