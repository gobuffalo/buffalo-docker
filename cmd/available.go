package cmd

import (
	"github.com/gobuffalo/buffalo-docker/docker"
	"github.com/gobuffalo/buffalo-plugins/plugins/plugcmds"
)

var Available = plugcmds.NewAvailable()

func init() {
	Available.Add("root", dockerCmd)
	Available.Listen(docker.Listen)
	Available.Add("generate", generateCmd)
	Available.Mount(rootCmd)
}
