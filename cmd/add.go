package cmd

import (
	"cli-task-manager/db"
	"fmt"
	"github.com/logrusorgru/aurora/v4"
	"github.com/spf13/cobra"
	"strings"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task",
	Long:  `Add a new task to your todo reminder`,
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		err := db.AddToBucket("tasks", task)

		if err != nil {
			fmt.Println(aurora.Red("Error occurred whilst adding task: " + task))
			return
		}

		fmt.Println("Successfully added \"" + task + "\" to task list")
	},
}
