package cmd

import (
	"fmt"
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
		fmt.Println("Done")
	},
}
