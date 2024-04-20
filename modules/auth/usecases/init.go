package usecases

import (
	ur "svc-user-management/modules/user/repositories"

	geninfra "svc-user-management/infrastructure/generator"
	jwtinfra "svc-user-management/infrastructure/jwt"
)

type Opts struct {
	UserRepositories ur.UserRepositories
	SnowflakeManager geninfra.SnowflakeManager
	JWTManager       jwtinfra.JWTManager
}

type Module struct {
	userRepositories ur.UserRepositories
	snowflakeManager geninfra.SnowflakeManager
	jWTManager       jwtinfra.JWTManager
}

func New(o *Opts) *Module {
	return &Module{
		userRepositories: o.UserRepositories,
		snowflakeManager: o.SnowflakeManager,
		jWTManager:       o.JWTManager,
	}
}
