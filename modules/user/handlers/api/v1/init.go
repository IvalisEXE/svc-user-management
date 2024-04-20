package v1

import (
	"svc-user-management/modules/user/usecases"
)

type Handler struct {
	userUsecases usecases.UserUsecases
}

func New(u usecases.UserUsecases) *Handler {
	return &Handler{u}
}
