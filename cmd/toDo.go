package cmd

import (
	"os"
	"task-tracker/utils"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

// toDoCmd represents the toDo command
var toDoCmd = &cobra.Command{
	Use:   "todo",
	Short: "Show all tasks with \"To Do\" status",
	Long:  `Show all tasks that are currently in the task tracker with the status "To Do".`,
	Run: func(cmd *cobra.Command, args []string) {
		listTasksToDo()
	},
}

func init() {
	listCmd.AddCommand(toDoCmd)
}

func listTasksToDo() {
	task, err := utils.LoadTasks("To do")

	if err != nil {
		return
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Description", "Status", "Created At", "Updated At"})

	for _, task := range task {
		t.AppendRow([]interface{}{task.Id, task.Description, task.Status, task.CreatedAt, task.UpdatedAt})
	}

	t.Render()
}
