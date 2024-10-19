package cmd

import (
	"os"
	"task-tracker/utils"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "Show all tasks with \"Done\" status",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		listTasksDone()
	},
}

func init() {
	listCmd.AddCommand(doneCmd)
}

func listTasksDone() {
	task, err := utils.LoadTasks("Done")

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
