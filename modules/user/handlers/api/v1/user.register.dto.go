package v1

import (
	"errors"
	"regexp"
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation"

	constant "svc-user-management/const"
	"svc-user-management/lib/common"
	"svc-user-management/modules/user/models"
)

type UserRegisterBodyReq struct {
	PhoneNo  string `json:"phone_no"`
	FullName string `json:"fullname"`
	Password string `json:"password"`
}

func (ur *UserRegisterBodyReq) ParseToRegisterUserEntity() models.User {
	return models.User{
		PhoneNo:  ur.PhoneNo,
		FullName: ur.FullName,
		Password: []byte(ur.Password),
	}
}

func (ur *UserRegisterBodyReq) NormalizeMSISDN() error {
	phoneno, err := common.NormalizeMSISDN(ur.PhoneNo)
	if err != nil {
		return errors.New("failed normalize phone number")
	}

	ur.PhoneNo = phoneno
	return nil
}

func (ur *UserRegisterBodyReq) RegiterUserValidation() error {

	if err := ur.LenghtValidation(); err != nil {
		return err
	}

	if err := ur.FormatPhoneNoValidation(); err != nil {
		return err
	}

	if err := ur.FormatCapitalPasswordValidation(); err != nil {
		return err
	}

	if err := ur.FormatAlphaNumericPasswordValidation(); err != nil {
		return err
	}

	if err := ur.FormatNonAlphaNumericPasswordValidation(); err != nil {
		return err
	}

	return nil
}

func (ur *UserRegisterBodyReq) LenghtValidation() error {
	if err := validation.Validate(ur.PhoneNo,
		validation.Required,
		validation.Length(8, 13),
	); err != nil {
		return errors.New("phone number must be have min 8 characters and max 13 characters")
	}
	if err := validation.Validate(ur.Password,
		validation.Required,
		validation.Length(6, 64),
	); err != nil {
		return errors.New("password must be have min 6 characters and max 64 characters")
	}
	return nil
}

func (ur *UserRegisterBodyReq) FormatPhoneNoValidation() error {
	phoneNoOk, err := regexp.MatchString(constant.REGEX_USER_PHONE_NO_VALIDATION, ur.PhoneNo)
	if err != nil {
		return errors.New("failed validation format phone number user register")
	}

	if !phoneNoOk {
		return errors.New("phone number must be have +62 firts")
	}

	return nil
}

func (ur *UserRegisterBodyReq) FormatCapitalPasswordValidation() error {
	reg := regexp.MustCompile(constant.REGEX_USER_PASSWORD_CAPITAL)
	capOk := reg.FindString(ur.Password) != ""
	if !capOk {
		return errors.New("password must be have 1 capital characters")
	}

	return nil
}

func (ur *UserRegisterBodyReq) FormatAlphaNumericPasswordValidation() error {

	alphaNumericOk, err := regexp.MatchString(constant.REGEX_USER_PASSWORD_ALPHA_NUMERIC, ur.Password)
	if err != nil {
		return errors.New("failed validation alpha numeric format passsword user register")
	}

	if !alphaNumericOk {
		return errors.New("password must be have 1 alpha numeric")
	}

	return nil
}

func (ur *UserRegisterBodyReq) FormatNonAlphaNumericPasswordValidation() error {

	reg := regexp.MustCompile(constant.REGEX_USER_PASSWORD_NON_ALPHA_NUMERIC)
	nonAlphaNumericOk := reg.MatchString(ur.Password)
	if nonAlphaNumericOk {
		return errors.New("password must be dont have non alpha numeric")
	}

	return nil
}

type UserRegisterBodyRes struct {
	UserRegisterUserBodyRes UserRegisterUserBodyRes `json:"user"`
}

type UserRegisterUserBodyRes struct {
	UserID string `json:"user_id"`
}

func (ur *UserRegisterBodyRes) ToRegisterUserResponse(userID int64) {
	ur.UserRegisterUserBodyRes.UserID = strconv.FormatInt(userID, 10)
}

type UserProfileBodyRes struct {
	UserProfileUserBodyRes UserProfileUserBodyRes `json:"user"`
}

type UserProfileUserBodyRes struct {
	PhoneNo  string `json:"phone_no"`
	Fullname string `json:"fullname"`
}
