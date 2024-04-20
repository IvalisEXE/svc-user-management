package repositories

import (
	"context"
	"fmt"
)

func (m *Module) UpsertUserLogin(ctx context.Context, userLogin UserLogin) error {
	var count int64
	row := m.pgm.Architecture().GetMaster().QueryRowxContext(ctx, findCountUserLoginByUserID, userLogin.UserID)
	row.Scan(&count)

	if count > 0 {
		_, err := m.pgm.Architecture().GetMaster().ExecContext(ctx, updateUserLogin,
			userLogin.LastLogin,
			userLogin.UserID,
		)
		if err != nil {
			fmt.Println("======= update", err.Error())
			return err
		}
	} else {
		_, err := m.pgm.Architecture().GetMaster().ExecContext(ctx, insertUserLogin,
			userLogin.UserID,
			1,
			userLogin.LastLogin,
			userLogin.CreatedBy,
			userLogin.CreatedAt,
		)

		if err != nil {
			fmt.Println("======= insert", err.Error())
			return err
		}
	}

	return nil
}
