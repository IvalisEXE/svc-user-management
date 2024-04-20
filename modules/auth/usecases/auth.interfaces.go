package usecases

import (
	"context"
)

type AuthUsecases interface {
	Login(ctx context.Context, phone, password string) (string, int64, error)
}
