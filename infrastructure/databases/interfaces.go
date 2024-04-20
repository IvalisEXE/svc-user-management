package databases

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"

	pgdb "svc-user-management/lib/databases"
)

type ExecContext func(ctx context.Context, query string, args ...interface{}) (sql.Result, error)

type GetContext func(ctx context.Context, dest interface{}, query string, args ...interface{}) error

type SelectContext func(ctx context.Context, dest interface{}, query string, args ...interface{}) error

//go:generate mockery --name DatabaseManagerPostgresql
type PGManager interface {
	Architecture() *pgdb.Architecture
	StartTransaction(ctx context.Context) (*sqlx.Tx, error)
	GetTransaction(dbTx *sqlx.Tx) ExecContext
	GetMaster(ctx context.Context, dbTx *sqlx.Tx) GetContext
	SelectMaster(ctx context.Context, dbTx *sqlx.Tx) SelectContext
}
