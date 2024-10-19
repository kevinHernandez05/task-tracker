package cmd

import (
	"os"
	"task-tracker/utils"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

// inProgressCmd represents the inProgress command
var inProgressCmd = &cobra.Command{
	Use:   "in-progress",
	Short: "Show all tasks with \"In Progress\" status",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		listTasksInProgress()
	},
}

func init() {
	listCmd.AddCommand(inProgressCmd)
}

func listTasksInProgress() {
	task, err := utils.LoadTasks("In Progress")

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
