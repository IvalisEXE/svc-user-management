package auth

import "github.com/golang-jwt/jwt/v4"

type SessionUser struct {
	UserID   string `json:"user_id"`
	PhoneNo  string `json:"phone_no"`
	FullName string `json:"fullname"`
}

type CustomClaims struct {
	UserID   string `json:"user_id"`
	PhoneNo  string `json:"phone_no"`
	Fullname string `json:"fullname"`
	jwt.RegisteredClaims
}
