package jwt

import (
	"svc-user-management/lib/auth"
)

func (m *Module) Generate(sessionUser auth.SessionUser) (string, error) {
	return m.aj.Generate(sessionUser)
}

func (m *Module) Parse(token string) (auth.SessionUser, error) {
	return m.aj.Parse(token)
}

func (m *Module) GetClaims(token string) (interface{}, error) {
	return m.aj.GetClaims(token)
}
