package initiator

import (
	"svc-user-management/config"
	"svc-user-management/init/register"

	ac "svc-user-management/modules/auth/usecases"
	uc "svc-user-management/modules/user/usecases"
)

var (
	userUsecases uc.UserUsecases
	authUsecases ac.AuthUsecases
)

func (i *Initiator) InitUsecases(cfg *config.Configuration, core *register.Core, infra *register.Infrastructure) *register.Usecases {
	userUsecases = uc.New(&uc.Opts{
		UserRepositories: userRepositories,
		SnowflakeManager: infra.SnowflakeManager,
	})
	authUsecases = ac.New(&ac.Opts{
		UserRepositories: userRepositories,
		SnowflakeManager: infra.SnowflakeManager,
		JWTManager:       infra.JWTManager,
	})
	return &register.Usecases{
		UserUsecases: userUsecases,
		AuthUsecases: authUsecases,
	}
}
