package main

import (
	"fmt"
	"path/filepath"
	"os"

	"cmd/cmd"
	"db/db"

	homedir "github.com/mitchellh/go-homedir"
)

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "todo.db")
	
	err := db.Init()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	cmd.RootCmd.Execute()
}
