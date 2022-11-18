package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "task",
	Short: "task is task manager for managing your todos",
	Long: `A Fast and Flexible Task To-do manager
			that helps you keep all your needs on your PC locally.
			Notifications coming soon!`,
	//Run: func(cmd *cobra.Command, args []string) {
	//	// Do Stuff Here
	//},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
