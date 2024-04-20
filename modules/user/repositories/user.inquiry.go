package repositories

import (
	"context"
	"database/sql"
)

func (m *Module) FindUserByID(ctx context.Context, userID int64) (FindUserByID, error) {
	var user FindUserByID
	if err := m.pgm.Architecture().GetMaster().GetContext(ctx, &user, findUserByID, userID); err != nil {
		if err == sql.ErrNoRows {
			return user, sql.ErrNoRows
		}
		return user, err
	}

	return user, nil
}

func (m *Module) FindCountUserByID(ctx context.Context, userID int64) (bool, error) {
	var count int64
	if err := m.pgm.Architecture().GetMaster().GetContext(ctx, &count, findCountUserByID, userID); err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return count > 0, nil
}

func (m *Module) FindUserByPhoneNo(ctx context.Context, phoneNo string) (FindUserByPhoneNo, error) {
	var user FindUserByPhoneNo
	row := m.pgm.Architecture().GetMaster().QueryRowxContext(ctx, findUserByPhoneNo, phoneNo)
	row.Scan(
		&user.UserID,
		&user.PhoneNo,
		&user.FullName,
		&user.Password,
		&user.CreatedAt,
	)

	return user, nil
}

func (m *Module) FindCountUserByPhoneNo(ctx context.Context, phoneNo string) (bool, error) {
	var count int64
	if err := m.pgm.Architecture().GetMaster().GetContext(ctx, &count, findCountUserByPhoneNo, phoneNo); err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return count > 0, nil
}
