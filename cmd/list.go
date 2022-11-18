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

		for k, v := range items {
			fmt.Printf("%s. %s\n", aurora.Magenta(k), v)
		}
	},
}
