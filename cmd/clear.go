package cmd

import (

	"github.com/spf13/cobra"
)

var clearCmd = &cobra.Command {
	Use: "clear",
	Short: "Clear all tasks from your todo list",
	Run: func(cmd *cobra.Command, args []strings) {
		//Clear all tasks from DB
	}
}

func init() {
	RootCmd.AddCommand(clearCmd)
}
