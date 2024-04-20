package repositories

import (
	"context"
)

func (m *Module) CreateUser(ctx context.Context, user CreateUser) (int64, error) {

	_, err := m.pgm.Architecture().GetMaster().ExecContext(ctx, createUser,
		user.UserID,
		user.PhoneNo,
		user.FullName,
		user.Password,
		user.CreatedAt,
		user.CreatedBy,
	)
	if err != nil {
		return 0, err
	}

	return user.UserID, nil
}
