package databases

import (
	"context"

	"github.com/jmoiron/sqlx"

	pgdb "svc-user-management/lib/databases"
)

func (m *Module) Architecture() *pgdb.Architecture {
	return m.db
}

func (m *Module) StartTransaction(ctx context.Context) (*sqlx.Tx, error) {
	/*
		init tracer
	*/

	return m.db.GetMaster().BeginTxx(ctx, nil)
}

func (m *Module) GetTransaction(dbTx *sqlx.Tx) ExecContext {
	/*
		init tracer
	*/

	if dbTx != nil {
		return dbTx.ExecContext
	}
	return m.Architecture().Master.ExecContext
}

func (m *Module) GetMaster(ctx context.Context, dbTx *sqlx.Tx) GetContext {
	/*
		init tracer
	*/

	if dbTx != nil {
		return dbTx.GetContext
	}

	return m.Architecture().GetMaster().GetContext
}

func (m *Module) SelectMaster(ctx context.Context, dbTx *sqlx.Tx) SelectContext {
	/*
		init tracer
	*/

	if dbTx != nil {
		return dbTx.SelectContext
	}

	return m.Architecture().GetMaster().SelectContext
}
