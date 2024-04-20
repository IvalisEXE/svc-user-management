package auth

import (
	"encoding/base64"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"

	constant "svc-user-management/const"
)

func (m *Module) Generate(auth SessionUser) (string, error) {
	authID := base64.StdEncoding.EncodeToString([]byte(time.Now().Format(constant.FORMAT_ISO_DATE)))
	authExpire := jwt.NewNumericDate(time.Now().Add(time.Duration(m.ttl * int(time.Second))))
	authClaims := CustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    auth.PhoneNo,
			ExpiresAt: authExpire,
			ID:        authID,
		},
		UserID:   auth.UserID,
		PhoneNo:  auth.PhoneNo,
		Fullname: auth.FullName,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, authClaims)
	tokenString, err := token.SignedString([]byte(m.signatureKey))
	if err != nil {
		return tokenString, err
	}

	return tokenString, err
}

func (m *Module) Parse(tokenString string) (SessionUser, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(m.signatureKey), nil
	})

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return SessionUser{
			UserID:   claims.UserID,
			PhoneNo:  claims.PhoneNo,
			FullName: claims.Fullname,
		}, nil
	}

	return SessionUser{}, err
}

func (m *Module) GetClaims(tokenString string) (interface{}, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(m.signatureKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("cannot parse claim")
}

// func (m *Module) CreateSessionKey(user interface{}, prefix string) string {

// 	claims, err := m.GetClaims(GetToken())
// 	var email string

// 	if err != nil {
// 		log.Println(err.Error())
// 		log.Println(map[string]interface{}{
// 			"data": user,
// 		})
// 	}

// 	if entity, ok := user.(entities.Auth); ok {

// 		email = entity.Email

// 	} else if entity, ok := user.(entities.SessionPayload); ok {

// 		email = entity.Email

// 	}

// 	return fmt.Sprintf(prefix+":%s&%s", email, claims.(*entities.CustomClaims).ID)
// }
