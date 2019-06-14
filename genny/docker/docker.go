package docker

import (
	"text/template"

	"github.com/gobuffalo/genny"
	"github.com/gobuffalo/gogen"
	"github.com/gobuffalo/packr/v2"
)

func New(opts *Options) (*genny.Generator, error) {
	g := genny.New()

	if err := opts.Validate(); err != nil {
		return g, err
	}

	data := map[string]interface{}{
		"opts": opts,
	}

	g.Box(packr.New("github.com/gobuffalo/buffalo-docker/common", "../docker/templates/common"))

	switch opts.Style {
	case "multi":
		g.Box(packr.New("github.com/gobuffalo/buffalo-docker/multi", "../docker/templates/multi"))
	case "standard":
		g.Box(packr.New("github.com/gobuffalo/buffalo-docker/standard", "../docker/templates/standard"))
	}

	helpers := template.FuncMap{}
	t := gogen.TemplateTransformer(data, helpers)
	g.Transformer(t)
	g.Transformer(genny.Dot())

	return g, nil
}
