package cmd

import (
	"cli-task-manager/db"
	"fmt"
	"github.com/logrusorgru/aurora/v4"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all todos",
	Long: `You shouldn't remember all your todos all the time, this 
			gives you all`,
	Run: func(cmd *cobra.Command, args []string) {
		items := db.ListBucketItems("tasks")
		current := 1

		for _, v := range items {
			fmt.Printf("%d. %s\n", aurora.Magenta(current), v)
			current++
		}
	},
}

//todo: Display something when map is empty
//todo: Display this after a task is done...
