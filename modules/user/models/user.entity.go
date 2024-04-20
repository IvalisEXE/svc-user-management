package models

import "time"

type User struct {
	UserID    int64
	PhoneNo   string
	FullName  string
	Password  []byte
	CreatedBy int64
	CreatedAt *time.Time
	UpdatedBy int64
	UpdatedAt *time.Time
}

type Users []User
