package cmd

import (
	"os"
	"task-tracker/utils"

	"github.com/jedib0t/go-pretty/v6/table"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Long:  `List all tasks that are currently in the task tracker.`,
	Run: func(cmd *cobra.Command, args []string) {
		listTasks()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func listTasks() {
	task, err := utils.LoadTasks("")

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
