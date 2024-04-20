package generator

import (
	"svc-user-management/lib/generator"
)

type Module struct {
	snode *generator.SnowflakeNode
}

type Opts struct {
	Snode *generator.SnowflakeNode
}

func New(o *Opts) *Module {
	return &Module{
		snode: o.Snode,
	}
}
