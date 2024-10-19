package cmd

import (
	"fmt"
	"strconv"
	"task-tracker/utils"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a task",
	Long:  `Update a task that is currently in the task tracker.`,
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])

		if err != nil {
			fmt.Println("The id must be a number")
			return
		}

		updateATask(args[1], id)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}

func updateATask(Description string, Id int) {

	// Cargar las tareas actuales
	tasks, err := utils.LoadTasks("")

	if err != nil {
		fmt.Println("Error loading tasks.", err)
		return
	}

	tasksWithTheUpdatedElement, index := utils.UpdateTaskById(Description, Id, tasks)

	if index == -1 {
		fmt.Println("Task not found")
		return
	}

	err = utils.SaveTasks(tasksWithTheUpdatedElement)
	fmt.Println("Task updated successfully")

	if err != nil {
		return
	}

}
