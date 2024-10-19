package cmd

import (
	"fmt"
	"strconv"
	"task-tracker/utils"

	"github.com/spf13/cobra"
)

// markInProgressCmd represents the markInProgress command
var markInProgressCmd = &cobra.Command{
	Use:   "mark-in-progress",
	Short: "Mark a task as in progress",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])

		if err != nil {
			fmt.Println("The id must be a number")
			return
		}

		updateStatusTaskAsInProgress(id)
	},
}

func init() {
	rootCmd.AddCommand(markInProgressCmd)
}

func updateStatusTaskAsInProgress(Id int) {

	// Cargar las tareas actuales
	tasks, err := utils.LoadTasks("")

	if err != nil {
		fmt.Println("Error loading tasks.", err)
		return
	}

	tasksWithTheUpdatedElement, index := utils.UpdateTaskStatusById("In Progress", Id, tasks)

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
