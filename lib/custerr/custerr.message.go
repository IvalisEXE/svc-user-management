package custerr

import "errors"

var (
	ErrMissingKeyReqHeader = errors.New("missing key in request header")
	ErrUnauthorized        = errors.New("Unauthorized")
	ErrInvalidCredentials  = errors.New("invalid credentials")
	ErrUserNotFound        = errors.New("user not found")
	ErrNoRows              = errors.New("sql: no rows in result set")
)
