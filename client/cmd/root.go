package cmd

import (
	//"fmt"
	//"os"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "cobra-cli",
		Short: "Store client app",
		Long:  "client app to manage file on a http server",
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
