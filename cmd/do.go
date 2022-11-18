package cmd

import (
	"cli-task-manager/db"
	"fmt"
	"github.com/logrusorgru/aurora/v4"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(doCmd)
}

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks a todo as complete",
	Long:  `Every todo has to be done after a while. This marks yours as done`,
	Run: func(cmd *cobra.Command, args []string) {
		err := db.RemoveFromBucket("tasks", args[0])
		if err != nil {
			fmt.Printf("Error occured: %s", aurora.Red(err))
		}
	},
}
