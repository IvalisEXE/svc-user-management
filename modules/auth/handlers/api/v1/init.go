package v1

import (
	ac "svc-user-management/modules/auth/usecases"
	uc "svc-user-management/modules/user/usecases"
)

type Handler struct {
	authUsecases ac.AuthUsecases
	userUsecases uc.UserUsecases
}

func New(auth ac.AuthUsecases, u uc.UserUsecases) *Handler {
	return &Handler{
		authUsecases: auth,
		userUsecases: u,
	}
}
