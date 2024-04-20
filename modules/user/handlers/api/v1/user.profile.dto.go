package v1

import (
	"errors"
	"regexp"
	constant "svc-user-management/const"
	"svc-user-management/lib/common"

	validation "github.com/go-ozzo/ozzo-validation"
)

type UpdateUserProfileBodyReq struct {
	PhoneNo  string `json:"phone_no"`
	FullName string `json:"fullname"`
}

func (uu *UpdateUserProfileBodyReq) UpdateUserProfileValidation() error {
	if err := uu.LenghtValidation(); err != nil {
		return err
	}

	if err := uu.FormatPhoneNoValidation(); err != nil {
		return err
	}

	return uu.NormalizeMSISDN()
}

func (uu *UpdateUserProfileBodyReq) LenghtValidation() error {
	if err := validation.Validate(uu.PhoneNo,
		validation.Required,
		validation.Length(8, 13),
	); err != nil {
		return errors.New("phone number must be have min 8 characters and max 13 characters")
	}

	return nil
}

func (uu *UpdateUserProfileBodyReq) FormatPhoneNoValidation() error {
	phoneNoOk, err := regexp.MatchString(constant.REGEX_USER_PHONE_NO_VALIDATION, uu.PhoneNo)
	if err != nil {
		return errors.New("failed validation format phone number user register")
	}

	if !phoneNoOk {
		return errors.New("phone number must be have +62 firts")
	}

	return nil
}

func (uu *UpdateUserProfileBodyReq) NormalizeMSISDN() error {
	phoneno, err := common.NormalizeMSISDN(uu.PhoneNo)
	if err != nil {
		return errors.New("failed normalize phone number")
	}

	uu.PhoneNo = phoneno
	return nil
}
