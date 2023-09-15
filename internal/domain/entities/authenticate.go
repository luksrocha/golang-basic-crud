package entities

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/luksrocha/house-system/constants"
	"github.com/luksrocha/house-system/util"
	"github.com/spf13/viper"
)

var (
	JWT_SECRET_KEY           = viper.GetString(constants.EnvKeyConstants().JWTSecretKey)
	EXPIRATION_TIME_ONE_WEEK = time.Hour * 24 * 7
)

func (u *User) GenerateToken() (string, error) {
	secret := JWT_SECRET_KEY

	claims := util.NewClaims(time.Now().Add(EXPIRATION_TIME_ONE_WEEK).Unix(), u.ID.String(), u.GetFullName(), u.Email)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", errors.New("error while signing the token")
	}

	teste, _ := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(token *jwt.Token) (any, error) {
		return token, nil
	})

	fmt.Println(teste.Claims)

	return tokenString, nil
}
