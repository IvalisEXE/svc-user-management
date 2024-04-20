package initiator

import (
	"svc-user-management/config"
	"svc-user-management/init/register"
)

type InitiatorManager interface {
	InitConfig() *config.Configuration
	InitCore(cfg *config.Configuration) *register.Core
	InitInfrastructure(cfg *config.Configuration, core *register.Core) *register.Infrastructure
	InitUsecases(cfg *config.Configuration, core *register.Core, infra *register.Infrastructure) *register.Usecases
	InitRepositories(cfg *config.Configuration, core *register.Core, infra *register.Infrastructure) *register.Repositories
}

type Initiator struct {
}

func New() InitiatorManager {
	return &Initiator{}
}
