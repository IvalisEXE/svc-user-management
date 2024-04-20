package repositories

import "context"

func (m *Module) UpdateUserProfileByID(ctx context.Context, userID int64, phoneNo, fullname string) error {
	_, err := m.pgm.Architecture().GetMaster().ExecContext(ctx, updateUserProfileByID,
		phoneNo,
		fullname,
		userID,
	)
	if err != nil {
		return err
	}
	return nil
}

func (m *Module) UpdateUserFullnameByID(ctx context.Context, userID int64, phoneNo, fullname string) error {
	_, err := m.pgm.Architecture().GetMaster().ExecContext(ctx, updateUserFullnameByID,
		fullname,
		userID,
	)
	if err != nil {
		return err
	}
	return nil
}
