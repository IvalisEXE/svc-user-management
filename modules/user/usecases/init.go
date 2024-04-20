package usecases

import (
	ur "svc-user-management/modules/user/repositories"

	geninfra "svc-user-management/infrastructure/generator"
)

type Opts struct {
	UserRepositories ur.UserRepositories
	SnowflakeManager geninfra.SnowflakeManager
}

type Module struct {
	userRepositories ur.UserRepositories
	snowflakeManager geninfra.SnowflakeManager
}

func New(o *Opts) *Module {
	return &Module{
		userRepositories: o.UserRepositories,
		snowflakeManager: o.SnowflakeManager,
	}
}
