package cmd

import (
	"eg-performance/helloworld"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:   "hello-word",
	Short: "Test Hello Word",
	Run: func(cmd *cobra.Command, args []string) {
		helloworld.RunHelloWord()
	},
}
