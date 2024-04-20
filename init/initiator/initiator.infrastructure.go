package initiator

import (
	"log"
	"svc-user-management/config"
	"svc-user-management/init/register"

	authLib "svc-user-management/lib/auth"
	genLib "svc-user-management/lib/generator"

	pginfra "svc-user-management/infrastructure/databases"
	geninfra "svc-user-management/infrastructure/generator"
	jwtinfra "svc-user-management/infrastructure/jwt"
)

var (
	pgManager        pginfra.PGManager
	snowflakeManager geninfra.SnowflakeManager
	jWTManager       jwtinfra.JWTManager
)

func (i *Initiator) InitInfrastructure(cfg *config.Configuration, core *register.Core) *register.Infrastructure {
	i.initPostgresSQL(cfg, core)
	i.initSnowflake(cfg, core)
	i.initJWT(cfg, core)
	return &register.Infrastructure{
		PGManagerr:       pgManager,
		SnowflakeManager: snowflakeManager,
		JWTManager:       jWTManager,
	}
}

func (i *Initiator) initPostgresSQL(cfg *config.Configuration, core *register.Core) {
	pgManager = pginfra.New(&pginfra.Opts{
		DB: core.PGArchitecture,
	})
}

func (i *Initiator) initSnowflake(cfg *config.Configuration, core *register.Core) {
	snode, err := genLib.New(1)
	if err != nil {
		log.Fatalf("Fatal init snowflake %s", err.Error())
	}

	snowflakeManager = geninfra.New(&geninfra.Opts{
		Snode: snode,
	})
}

func (i *Initiator) initJWT(cfg *config.Configuration, core *register.Core) {
	authJWT := authLib.New(&authLib.Opts{
		TTL:          cfg.JWT.TTL,
		SignatureKey: cfg.JWT.SignatureKey,
	})

	jWTManager = jwtinfra.New(&jwtinfra.Opts{
		Aj: authJWT,
	})
}
