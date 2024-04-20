package api

import (
	"github.com/labstack/echo/v4"

	"svc-user-management/config"
	"svc-user-management/init/register"
	mymiddleware "svc-user-management/middleware"
	v1 "svc-user-management/modules/user/handlers/api/v1"
)

func Register(e *echo.Echo, uc *register.Usecases, cfg *config.Configuration) {
	h := v1.New(uc.UserUsecases)

	routeV1 := e.Group("api/v1/user")
	routeV1.POST("", h.RegisterUser)

	routeV1.Use(mymiddleware.JwtMiddleware(cfg.JWT.SignatureKey))
	routeV1.GET("", h.GetProfileUser)
	routeV1.PUT("", h.UpdateProfileUser)
}
