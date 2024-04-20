package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"

	"svc-user-management/lib/response"
)

type UserAuth struct {
	Phoneno  string
	Fullname string
}

// Custom JWT middleware
func JwtMiddleware(tokenKeys ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) (err error) {
			// Get token from header
			authorizationHeader := ctx.Request().Header.Get("Authorization")

			// Check if string has token
			if !strings.Contains(authorizationHeader, "Bearer") {
				response := map[string]interface{}{
					"status": map[string]interface{}{
						"code":    response.StatusCodeUnauthorized,
						"message": "invalid token or missing token",
						"errors":  "invalid token or missing token",
					},
				}
				return ctx.JSON(http.StatusUnauthorized, response)
			}

			var (
				token       *jwt.Token
				tokenString = strings.Replace(authorizationHeader, "Bearer ", "", -1)
			)

			for _, key := range tokenKeys {
				if key == "" {
					continue
				}
				// Parse token
				token, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
					if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok || method != jwt.SigningMethodHS256 {
						return nil, fmt.Errorf("signing method invalid")
					}

					return []byte(key), nil
				})

				if err == nil {
					break
				}
			}

			// Check if there is error
			if err != nil {
				response := map[string]interface{}{
					"status": map[string]interface{}{
						"code":    response.StatusCodeUnauthorized,
						"message": "something went wrong when parse token",
						"errors":  err.Error(),
					},
				}
				return ctx.JSON(http.StatusUnauthorized, response)
			}

			// Save user_id and branch_id to context
			claims := token.Claims.(jwt.MapClaims)

			// Set to Context
			ctx.Set("user_id", claims["user_id"])
			ctx.Set("phone_no", claims["phone_no"])
			ctx.Set("fullname", claims["fullname"])

			// Skip to router
			return next(ctx)
		}

	}
}
