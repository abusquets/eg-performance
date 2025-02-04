package cmd

import (
	"eg-performance/eg"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(apiCmd)
}

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Test Api",
	Run: func(cmd *cobra.Command, args []string) {
		eg.RunApi()
	},
}
