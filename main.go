package main

import (
	"svc-user-management/cmd"

	_ "svc-user-management/docs"
)

// @title User Management - API
// @version 1.0
// @description Simple User Management Service
// @host localhost:8080
// @BasePath /

// @securityDefinitions.basic  BasicAuth

// @securityDefinitions.apikey XForwardedForAuth
// @in header
// @name X-Forwarded-For
func main() {
	cmd.Execute()
}
