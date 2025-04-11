package cmd

import (
	"strings"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command {
	Use: "do",
	Short: "Complete an item on your todo list",
	Run: func(cmd *cobra.Command, args []string) {
		var tasks []int
		for _, strArg := args {
			taskInt, err = strconv.Atoi(strArg)
			if err != nil{
				fmt.Printf("Argument %v invalid - must be an integer value.", strArg)
				continue
			}
			tasks = append(tasks, taskInt)
		}
	}
}

func init(){
	RootCmd.AddCommand(doCmd)
} 
