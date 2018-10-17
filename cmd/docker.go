package cmd

import (
	"github.com/spf13/cobra"
)

// dockerCmd represents the buffalo-docker command
var dockerCmd = &cobra.Command{
	Use:   "buffalo-docker",
	Short: "tools for working with docker",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	rootCmd.AddCommand(dockerCmd)
}
