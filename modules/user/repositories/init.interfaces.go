package repositories

import (
	"context"
)

type UserRepositories interface {
	CreateUser(ctx context.Context, user CreateUser) (int64, error)

	FindUserByID(ctx context.Context, userID int64) (FindUserByID, error)
	FindCountUserByID(ctx context.Context, userID int64) (bool, error)

	FindUserByPhoneNo(ctx context.Context, phoneNo string) (FindUserByPhoneNo, error)
	FindCountUserByPhoneNo(ctx context.Context, phoneNo string) (bool, error)

	UpsertUserLogin(ctx context.Context, user UserLogin) error

	UpdateUserProfileByID(ctx context.Context, userID int64, phoneNo, fullname string) error
	UpdateUserFullnameByID(ctx context.Context, userID int64, phoneNo, fullname string) error
}
