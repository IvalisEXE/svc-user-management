package auth

type AuthJWT interface {
	Generate(sessionUser SessionUser) (string, error)
	Parse(token string) (SessionUser, error)
	GetClaims(token string) (interface{}, error)
	// CreateSessionKey(user interface{}, prefix string) string
}
