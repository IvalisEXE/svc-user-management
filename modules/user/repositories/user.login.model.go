package repositories

import "time"

type UserLogin struct {
	ID         int64      `db:"id"`
	UserID     int64      `db:"user_id"`
	TotalLogin int        `db:"total_login"`
	LastLogin  *time.Time `db:"last_login"`
	CreatedBy  int64      `db:"created_by"`
	CreatedAt  *time.Time `db:"created_at"`
}
