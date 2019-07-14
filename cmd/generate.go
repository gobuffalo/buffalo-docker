package cmd

import (
	"context"

	"github.com/gobuffalo/buffalo-docker/genny/docker"
	"github.com/gobuffalo/genny"
	"github.com/gobuffalo/gogen"
	"github.com/spf13/cobra"
)

var generateOptions = struct {
	*docker.Options
	dryRun bool
}{
	Options: &docker.Options{},
}

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:     "docker",
	Short:   "generates a new docker",
	Aliases: []string{"dockerx"},
	RunE: func(cmd *cobra.Command, args []string) error {
		r := genny.WetRunner(context.Background())

		if generateOptions.dryRun {
			r = genny.DryRunner(context.Background())
		}

		g, err := docker.New(generateOptions.Options)
		if err != nil {
			return err
		}
		r.With(g)

		g, err = gogen.Fmt(r.Root)
		if err != nil {
			return err
		}
		r.With(g)

		return r.Run()
	},
}

func init() {
	generateCmd.Flags().StringVar(&generateOptions.Style, "style", "multi", "what style Dockerfile to generate [multi, standard]")
	generateCmd.Flags().BoolVarP(&generateOptions.dryRun, "dry-run", "d", false, "run the generator without creating files or running commands")
	dockerCmd.AddCommand(generateCmd)
}
