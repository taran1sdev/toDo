package cmd

import (
	"fmt"

	"todo/db"

	"github.com/spf13/cobra"
)

var clearCmd = &cobra.Command {
	Use: "clear",
	Short: "Clear all tasks from your todo list",
	Run: func(cmd *cobra.Command, args []string) {
		err := db.Clear()
		if err != nil {
			fmt.Println("Unable to clear your todo list!")
		} else {
			
			fmt.Println("Tasks cleared successfully")
		}
	},
}

func init() {
	RootCmd.AddCommand(clearCmd)
}
