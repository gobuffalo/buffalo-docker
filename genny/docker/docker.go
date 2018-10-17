package docker

import (
	"text/template"

	"github.com/gobuffalo/genny"
	"github.com/gobuffalo/genny/movinglater/gotools"
	"github.com/gobuffalo/packr"
	"github.com/pkg/errors"
)

func New(opts *Options) (*genny.Generator, error) {
	g := genny.New()

	if err := opts.Validate(); err != nil {
		return g, errors.WithStack(err)
	}

	data := map[string]interface{}{
		"opts": opts,
	}

	g.Box(packr.NewBox("../docker/templates/common"))

	switch opts.Style {
	case "multi":
		g.Box(packr.NewBox("../docker/templates/multi"))
	case "standard":
		g.Box(packr.NewBox("../docker/templates/standard"))
	}

	helpers := template.FuncMap{}
	t := gotools.TemplateTransformer(data, helpers)
	g.Transformer(t)
	g.Transformer(genny.Dot())

	return g, nil
}
