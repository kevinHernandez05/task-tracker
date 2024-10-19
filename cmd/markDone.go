package cmd

import (
	"fmt"
	"strconv"
	"task-tracker/utils"

	"github.com/spf13/cobra"
)

var markDoneCmd = &cobra.Command{
	Use:   "mark-done",
	Short: "Mark a task as done",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])

		if err != nil {
			fmt.Println("The id must be a number")
			return
		}

		updateStatusTaskAsDone(id)
	},
}

func init() {
	rootCmd.AddCommand(markDoneCmd)
}

func updateStatusTaskAsDone(Id int) {

	// Cargar las tareas actuales
	tasks, err := utils.LoadTasks("")

	if err != nil {
		fmt.Println("Error loading tasks.", err)
		return
	}

	tasksWithTheUpdatedElement, index := utils.UpdateTaskStatusById("Done", Id, tasks)

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
