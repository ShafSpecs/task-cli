package cmd

import (
	"fmt"
	"github.com/ShafSpecs/task/db"
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

		if len(items) == 0 {
			fmt.Println(aurora.Bold(aurora.Magenta("No task added yet!")), "Use the \"add\" command to start adding tasks.")
			return
		}

		current := 1

		for _, v := range items {
			fmt.Printf("%d. %s\n", aurora.Magenta(current), v)
			current++
		}
	},
}

//todo: Flags for version, etc.
//todo: config file
//todo: Display this after a task is done...
