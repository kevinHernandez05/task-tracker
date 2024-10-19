/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"
	"task-tracker/utils"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task",
	Long:  `Delete a task that is currently in the task tracker.`,
	Run: func(cmd *cobra.Command, args []string) {

		id, err := strconv.Atoi(args[0])

		if err != nil {
			fmt.Println("The id must be a number")
			return
		}

		deleteATask(id)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}

func deleteATask(Id int) {

	// Cargar las tareas actuales
	tasks, err := utils.LoadTasks("")

	if err != nil {
		fmt.Println("Error loading tasks.", err)
		return
	}

	tasksWithTheDeletedElement, index := utils.DeleteTaskById(tasks, Id)

	if index == -1 {
		fmt.Println("Task not found")
		return
	}

	err = utils.SaveTasks(tasksWithTheDeletedElement)
	fmt.Println("Task deleted successfully")

	if err != nil {
		return
	}

}
