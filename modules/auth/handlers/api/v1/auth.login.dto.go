package v1

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation"

	"svc-user-management/lib/common"
)

type Credential struct {
	PhoneNo  string `json:"phone_no"`
	Password string `json:"password"`
}

func (c *Credential) LengthValidation() error {
	if err := validation.Validate(c.PhoneNo,
		validation.Required,
		validation.Length(8, 13),
	); err != nil {
		return errors.New("phone number must be have min 8 characters and max 13 characters")
	}

	return nil
}

func (c *Credential) NormalizeMSISDN() error {
	phoneno, err := common.NormalizeMSISDN(c.PhoneNo)
	if err != nil {
		return errors.New("failed normalize phone number")
	}

	c.PhoneNo = phoneno
	return nil
}

type LoginBodyRes struct {
	User  LoginBodyUserRes `json:"user"`
	Token string           `json:"token"`
}

type LoginBodyUserRes struct {
	UserID string `json:"user_id"`
}
