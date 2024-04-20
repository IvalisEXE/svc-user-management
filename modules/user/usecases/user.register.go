package usecases

import (
	"context"
	"time"

	"svc-user-management/modules/user/models"
	"svc-user-management/modules/user/repositories"
)

func (m *Module) CheckUserExistByPhoneNo(ctx context.Context, phoneNo string) (bool, error) {
	return m.userRepositories.FindCountUserByPhoneNo(ctx, phoneNo)
}

func (m *Module) CreateUser(ctx context.Context, user models.User) (int64, error) {
	/*
		Need Checking User Eksist
	*/

	now := time.Now()
	uid := m.snowflakeManager.GenerateID()
	return m.userRepositories.CreateUser(ctx, repositories.CreateUser{
		UserID:    uid,
		FullName:  user.FullName,
		PhoneNo:   user.PhoneNo,
		Password:  user.Password,
		CreatedAt: &now,
		CreatedBy: uid,
	})
}
