package main

import (
	"fmt"
	"github.com/ShafSpecs/task/cmd"
	"github.com/ShafSpecs/task/db"
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
