package entities

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/luksrocha/house-system/constants"
	"github.com/spf13/viper"
)

var (
	JWT_SECRET_KEY           = viper.GetString(constants.EnvKeyConstants().JWTSecretKey)
	EXPIRATION_TIME_ONE_WEEK = time.Hour * 24 * 7
)

func (u *User) GenerateToken() (string, error) {
	secret := JWT_SECRET_KEY

	claims := jwt.MapClaims{
		"userId":   u.ID,
		"username": u.GetFullName(),
		"email":    u.Email,
		"exp":      time.Now().Add(EXPIRATION_TIME_ONE_WEEK).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", errors.New("error while signing the token")
	}

	return tokenString, nil
}
