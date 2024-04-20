package repositories

import (
	pgdb "svc-user-management/infrastructure/databases"
)

type Opts struct {
	PGM pgdb.PGManager
}

type Module struct {
	pgm pgdb.PGManager
}

func New(o *Opts) *Module {
	return &Module{
		pgm: o.PGM,
	}
}
