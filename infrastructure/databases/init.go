package databases

import (
	pgdb "svc-user-management/lib/databases"
)

type Module struct {
	db *pgdb.Architecture
}

type Opts struct {
	DB *pgdb.Architecture
}

func New(o *Opts) *Module {
	return &Module{
		db: o.DB,
	}
}
