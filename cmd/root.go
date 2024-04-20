package cmd

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"go.elastic.co/apm/module/apmechov4"

	"svc-user-management/init/assembler"

	api_auth "svc-user-management/modules/auth/handlers/api"
	api_user "svc-user-management/modules/user/handlers/api"
)

func Execute() {
	cwd, _ := os.Getwd()
	dirString := strings.Split(cwd, "svc-user-management")
	dir := strings.Join([]string{dirString[0], "svc-user-management"}, "")
	AppPath := dir

	godotenv.Load(filepath.Join(AppPath, "/.env"))

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(apmechov4.Middleware())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAuthorization,
		},
	}))

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	assembler := assembler.NewAssembler()
	assembler.BuildService()
	defer assembler.ReleaseResourcesAPI(e)

	api_user.Register(e, assembler.GetUsecases(), assembler.GetConfigutation())
	api_auth.Register(e, assembler.GetUsecases())

	go assembler.RunningAPI(e)
}
