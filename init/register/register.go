package register

import (
	pglib "svc-user-management/lib/databases"

	pginfra "svc-user-management/infrastructure/databases"
	geninfra "svc-user-management/infrastructure/generator"
	jwtinfra "svc-user-management/infrastructure/jwt"

	auth_usescases "svc-user-management/modules/auth/usecases"
	user_usescases "svc-user-management/modules/user/usecases"

	auth_repositories "svc-user-management/modules/auth/repositories"
	user_repositories "svc-user-management/modules/user/repositories"
)

type Core struct {
	PGArchitecture *pglib.Architecture
}

type Infrastructure struct {
	PGManagerr       pginfra.PGManager
	SnowflakeManager geninfra.SnowflakeManager
	JWTManager       jwtinfra.JWTManager
}

type Usecases struct {
	UserUsecases user_usescases.UserUsecases
	AuthUsecases auth_usescases.AuthUsecases
}

type Repositories struct {
	UserRepositories user_repositories.UserRepositories
	AuthRepositories auth_repositories.AuthRepositories
}
