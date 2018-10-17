package cmd

import (
	"fmt"

	"github.com/gobuffalo/buffalo-docker/docker"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "current version of docker",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("docker", docker.Version)
		return nil
	},
}

func init() {
	dockerCmd.AddCommand(versionCmd)
}
