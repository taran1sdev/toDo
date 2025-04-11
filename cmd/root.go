package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "todo",
	Short: "todo is a command-line task manager written in Go!",
	Long: `Usage:
		todo list: list all tasks
		todo add <task>: add task to list
		todo do <task> remove task from list`,
}
