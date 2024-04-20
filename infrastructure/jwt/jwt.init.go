package jwt

import (
	"svc-user-management/lib/auth"
)

type Opts struct {
	Aj auth.AuthJWT
}

type Module struct {
	aj auth.AuthJWT
}

func New(o *Opts) *Module {
	return &Module{
		aj: o.Aj,
	}
}
