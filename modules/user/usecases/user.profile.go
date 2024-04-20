package usecases

import (
	"context"

	"svc-user-management/modules/user/models"
)

func (m *Module) GetUserByID(ctx context.Context, userID int64) (*models.User, error) {
	user, err := m.userRepositories.FindUserByID(ctx, userID)
	if err != nil {
		return nil, nil
	}

	return &models.User{
		UserID:   user.UserID,
		PhoneNo:  user.PhoneNo,
		FullName: user.FullName,
	}, nil
}

func (m *Module) UpdateUserProfileByID(ctx context.Context, userID int64, phoneNo, fullname string, skipPhoneNo bool) error {
	if skipPhoneNo {
		return m.userRepositories.UpdateUserFullnameByID(ctx, userID, phoneNo, fullname)
	}
	return m.userRepositories.UpdateUserProfileByID(ctx, userID, phoneNo, fullname)
}
