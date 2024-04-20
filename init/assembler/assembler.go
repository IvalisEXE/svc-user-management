package assembler

import (
	"os"
	"os/signal"

	"github.com/labstack/echo/v4"

	"svc-user-management/config"
	"svc-user-management/init/initiator"
	"svc-user-management/init/register"
)

type AssemblerManager interface {
	BuildService() AssemblerManager

	RunningAPI(e *echo.Echo)
	ReleaseResourcesAPI(e *echo.Echo)

	GetConfigutation() *config.Configuration
	GetInfrastructure() *register.Infrastructure
	GetUsecases() *register.Usecases
	GetRepositories() *register.Repositories

	TerminateSignal()
}

type assembler struct {
	Initiator      initiator.InitiatorManager
	config         *config.Configuration
	core           *register.Core
	infrastructure *register.Infrastructure
	usecases       *register.Usecases
	repositories   *register.Repositories
}

func NewAssembler() AssemblerManager {
	return &assembler{
		Initiator: initiator.New(),
	}
}

func (a *assembler) BuildService() AssemblerManager {
	cfg := a.Initiator.InitConfig()
	core := a.Initiator.InitCore(cfg)
	infra := a.Initiator.InitInfrastructure(cfg, core)
	repositories := a.Initiator.InitRepositories(cfg, core, infra)
	usecases := a.Initiator.InitUsecases(cfg, core, infra)

	a.config = cfg
	a.core = core
	a.infrastructure = infra
	a.repositories = repositories
	a.usecases = usecases
	return a
}

func (a *assembler) GetInfrastructure() *register.Infrastructure {
	return a.infrastructure
}

func (a *assembler) GetUsecases() *register.Usecases {
	return a.usecases
}

func (a *assembler) GetRepositories() *register.Repositories {
	return a.repositories
}

func (a *assembler) GetConfigutation() *config.Configuration {
	return a.config
}

func (a *assembler) TerminateSignal() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
}
