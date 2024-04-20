package initiator

import (
	"svc-user-management/config"
	"svc-user-management/lib/env"
)

func (i *Initiator) InitConfig() *config.Configuration {
	return config.SetupConfiguration(env.NewConfiguration("svc-settlement"))
}
