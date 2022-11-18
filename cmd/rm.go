package cmd

import (
	"fmt"
	"github.com/ShafSpecs/task/db"
	"github.com/logrusorgru/aurora/v4"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(rmCmd)
}

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Deletes a todo task",
	Long:  `Deletes a task forever! No one would know...`,
	Run: func(cmd *cobra.Command, args []string) {
		err := db.DeleteFromBucket(args[0])
		if err != nil {
			fmt.Printf("Error occured: %s", aurora.Red(err))
		}
	},
}
