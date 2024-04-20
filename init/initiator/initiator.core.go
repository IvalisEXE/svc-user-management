package initiator

import (
	"svc-user-management/config"
	"svc-user-management/init/register"

	pglib "svc-user-management/lib/databases"
)

var (
	pdb *pglib.Architecture
)

func (i *Initiator) InitCore(cfg *config.Configuration) *register.Core {
	pdb = pglib.New(pglib.DBConfig{
		MasterDSN:     cfg.Database.DatabaseURL,
		MaxConn:       int(cfg.Database.DatabaseMaxCon),
		MaxIdleConn:   int(cfg.Database.DatabaseMaxIddleConn),
		RetryInterval: int(cfg.Database.DatabaseRetryInterval),
	}, pglib.DriverPostgres)

	return &register.Core{PGArchitecture: pdb}
}
