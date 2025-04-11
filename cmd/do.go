package cmd

import (
	"fmt"
	"strconv"
	"todo/db"
	
	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command {
	Use: "do",
	Short: "Complete an item on your todo list",
	Run: func(cmd *cobra.Command, args []string) {
		var tasks []int
		for _, strArg := range args {
			taskInt, err := strconv.Atoi(strArg)
			if err != nil{
				fmt.Printf("Argument %v invalid - must be an integer value.", strArg)
				continue
			}
			tasks = append(tasks, taskInt)
		}
		
		err := db.DoTasks(tasks)
		if err != nil {
			fmt.Println("Error occured while deleting task:", err)
		} else {
			newTasks, _ := db.GetTasks() 
			fmt.Println("Still left to do!")
			
			for _, task := range newTasks {
				fmt.Printf("%d) %v\n", task.Key, task.Value)
			}
		}
	    
	},
}

func init(){
	RootCmd.AddCommand(doCmd)
} 
