package cmd

import (
	"eg-performance/eg"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(frontendCmd)
}

var frontendCmd = &cobra.Command{
	Use:   "frontend",
	Short: "Test Fontend",
	Run: func(cmd *cobra.Command, args []string) {
		eg.RunFrontend()
	},
}
