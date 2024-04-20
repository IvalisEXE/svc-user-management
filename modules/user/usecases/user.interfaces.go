package usecases

import (
	"context"

	"svc-user-management/modules/user/models"
)

type UserUsecases interface {
	CheckUserExistByPhoneNo(ctx context.Context, phoneNo string) (bool, error)
	CreateUser(ctx context.Context, user models.User) (int64, error)

	GetUserByID(ctx context.Context, userID int64) (*models.User, error)

	UpdateUserProfileByID(ctx context.Context, userID int64, phoneNo, fullname string, skipPhoneNo bool) error
}
