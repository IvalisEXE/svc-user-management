package repositories

import "time"

type FindUserByID struct {
	UserID    int64      `db:"id"`
	PhoneNo   string     `db:"phone_no"`
	FullName  string     `db:"fullname"`
	Password  []byte     `db:"password"`
	CreatedAt *time.Time `db:"created_at"`
	CreatedBy int64      `db:"created_by"`
}

type FindUserByPhoneNo struct {
	UserID    int64      `db:"id"`
	PhoneNo   string     `db:"phone_no"`
	FullName  string     `db:"fullname"`
	Password  []byte     `db:"password"`
	CreatedAt *time.Time `db:"created_at"`
	CreatedBy int64      `db:"created_by"`
}
