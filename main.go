package main

import (
	"cli-task-manager/cmd"
	"cli-task-manager/db"
	"fmt"
	"github.com/logrusorgru/aurora/v4"
)

func main() {
	err := db.OpenDB()
	if err != nil {
		fmt.Println(aurora.Bold(aurora.Red("Task is already running!")))
		return
	}

	defer func() {
		err := db.CloseDB()
		if err != nil {
			fmt.Printf("Error closing memory file!\n"+
				"Error: %s", aurora.Red(err))
		}
	}()

	cmd.Execute()
}
