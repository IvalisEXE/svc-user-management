package config

import (
	"time"

	"svc-user-management/lib/env"
)

type Configuration struct {
	Database *Database
	JWT      *JWT
}

type (
	Database struct {
		DatabaseURL           string
		DatabaseRetryInterval time.Duration
		DatabaseMaxIddleConn  time.Duration
		DatabaseMaxCon        time.Duration
	}
	JWT struct {
		TTL          int
		SignatureKey string
	}
)

func SetupConfiguration(cfg env.Configuration) *Configuration {
	return &Configuration{
		Database: loadDatabase(cfg),
		JWT:      loadJWT(cfg),
	}
}

func loadDatabase(cfg env.Configuration) *Database {
	return &Database{
		DatabaseURL:           cfg.GetString("DATABASE_URL"),
		DatabaseRetryInterval: cfg.GetDuration("DATABASE_RETRY_INTERNAL"),
		DatabaseMaxIddleConn:  cfg.GetDuration("DATABASE_MAX_IDDLE_CONN"),
		DatabaseMaxCon:        cfg.GetDuration("DATABASE_MAX_CONN"),
	}
}

func loadJWT(cfg env.Configuration) *JWT {
	return &JWT{
		TTL:          cfg.GetInt("JWT_TTL"),
		SignatureKey: cfg.GetString("JWT_SIGNATURE_KEY"),
	}
}
