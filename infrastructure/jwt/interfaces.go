package jwt

import (
	"svc-user-management/lib/auth"
)

type JWTManager interface {
	Generate(sessionUser auth.SessionUser) (string, error)
	Parse(string) (auth.SessionUser, error)
	GetClaims(token string) (interface{}, error)
	// CreateSessionKey(user interface{}, prefix string) string
}
