package initiator

import (
	"svc-user-management/config"
	"svc-user-management/init/register"

	ar "svc-user-management/modules/auth/repositories"
	ur "svc-user-management/modules/user/repositories"
)

var (
	userRepositories ur.UserRepositories
	authRepositories ar.AuthRepositories
)

func (i *Initiator) InitRepositories(cfg *config.Configuration, core *register.Core, infra *register.Infrastructure) *register.Repositories {
	userRepositories = ur.New(&ur.Opts{
		PGM: infra.PGManagerr,
	})
	authRepositories = ar.New(&ar.Opts{
		PGM: infra.PGManagerr,
	})
	return &register.Repositories{
		UserRepositories: userRepositories,
		AuthRepositories: authRepositories,
	}
}
