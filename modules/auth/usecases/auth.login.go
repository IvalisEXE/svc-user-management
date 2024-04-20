package usecases

import (
	"context"
	"strconv"
	"time"

	"svc-user-management/lib/auth"
	"svc-user-management/lib/custerr"
	"svc-user-management/modules/user/repositories"
)

func (m *Module) Login(ctx context.Context, phoneNo, password string) (string, int64, error) {
	user, err := m.userRepositories.FindUserByPhoneNo(ctx, phoneNo)
	if err != nil {
		return "", user.UserID, err
	}

	passOK := string(user.Password) == password
	if !passOK {
		return "", user.UserID, custerr.ErrInvalidCredentials
	}

	token, err := m.jWTManager.Generate(auth.SessionUser{
		UserID:   strconv.FormatInt(user.UserID, 10),
		PhoneNo:  user.PhoneNo,
		FullName: user.FullName,
	})
	if err != nil {
		return token, user.UserID, err
	}

	now := time.Now()
	go m.userRepositories.UpsertUserLogin(context.Background(), repositories.UserLogin{
		UserID:    user.UserID,
		LastLogin: &now,
		CreatedBy: user.UserID,
		CreatedAt: &now,
	})

	return token, user.UserID, nil
}
