package cmd

import (
	"fmt"
	"github.com/ShafSpecs/task/db"
	"github.com/logrusorgru/aurora/v4"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(completedCmd)
}

var completedCmd = &cobra.Command{
	Use:   "completed",
	Short: "Shows completed tasks",
	Long:  `View all your completed tasks and give yourself a pat on the back`,
	Run: func(cmd *cobra.Command, args []string) {
		items := db.ListBucketItems("completed")

		fmt.Printf("You have finished the following tasks %s:\n", aurora.Magenta("today"))

		for _, v := range items {
			fmt.Printf("- %s\n", v)
		}
	},
}

//todo: Add ability to give `completed` timeframe
