package api

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"svc-user-management/init/register"
	v1 "svc-user-management/modules/auth/handlers/api/v1"
)

func Register(e *echo.Echo, uc *register.Usecases) {
	h := v1.New(uc.AuthUsecases, uc.UserUsecases)

	routeV1 := e.Group("api/v1/auth")
	routeV1.POST("/login", h.Login, middleware.BasicAuth(fetchCredential))
}

func fetchCredential(username, password string, c echo.Context) (bool, error) {
	cred := v1.Credential{
		PhoneNo:  username,
		Password: password,
	}

	c.Set("credential", cred)
	return true, nil
}
