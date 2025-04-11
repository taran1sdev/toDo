package cmd

import (
	"fmt"

	"todo/db"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use: "list",
	Short: "Print out your todo list",
	Run: func(cmd *cobra.Command, args []string){
		tasks, err := db.GetTasks()

		if err != nil {
			fmt.Println("Error reading database")
		} else if len(tasks) == 0 {
			fmt.Println("You have no tasks in your todo list!")
		} else {
			for _, task := range tasks {
				fmt.Printf("%d) %v\n", task.Key, task.Value)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
