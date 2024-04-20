package repositories

import "time"

type CreateUser struct {
	UserID    int64      `db:"id"`
	PhoneNo   string     `db:"phone_no"`
	FullName  string     `db:"fullname"`
	Password  []byte     `db:"password"`
	CreatedBy int64      `db:"created_at"`
	CreatedAt *time.Time `db:"created_by"`
}
