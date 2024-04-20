package repositories

import (
	"context"
)

type AuthRepositories interface {
	Login(ctx context.Context) error
}
