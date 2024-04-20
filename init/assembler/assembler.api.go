package assembler

import (
	"os"

	"github.com/labstack/echo/v4"
)

func (a *assembler) RunningAPI(e *echo.Echo) {
	if err := e.Start(":" + os.Getenv("SERVER_HTTP_PORT")); err != nil {
		e.Logger.Info("Shutting down the server")
	}
}
