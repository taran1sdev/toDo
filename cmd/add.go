package cmd

import (
	"fmt"
	"strings"
	
	"todo/db"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use: "add",
	Short: "Add a task to the todo list",
	Example: "todo add Finish writing README.md",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		id, err := db.AddTask(task)
		if err != nil {
			fmt.Println("Could not add task to db error:", err)
		} else {
			fmt.Printf("Task: %d) '%v' added to todo list.\n", id, task)
		}
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
