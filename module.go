package main

import (
	"github.com/lucasclerissepro/protoc-gen-temporal/templates/golang"
	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
	"text/template"
)

type TemporalModule struct {
	*pgs.ModuleBase
	ctx pgsgo.Context
}

func Temporal() pgs.Module {
	return &TemporalModule{ModuleBase: &pgs.ModuleBase{}}
}

func (m *TemporalModule) Name() string {
	return "temporal"
}

func (m *TemporalModule) InitContext(ctx pgs.BuildContext) {
	m.ModuleBase.InitContext(ctx)
	m.ctx = pgsgo.InitContext(ctx.Parameters())
}

func (m *TemporalModule) Execute(targets map[string]pgs.File, pkgs map[string]pgs.Package) []pgs.Artifact {
	tpl := template.Must(template.New("workflow").Parse(golang.WorkflowTpl))

	for _, f := range targets {
		outPath := m.ctx.OutputPath(f).SetExt(".temporal.go").String()
		m.AddGeneratorTemplateFile(outPath, tpl, f)
	}

	return m.Artifacts()
}
