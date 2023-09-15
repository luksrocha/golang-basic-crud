package util

import (
	"net/http"

	"github.com/golang-jwt/jwt"
)

func NewClaims(expirationTime int64, userId, username, email string) jwt.MapClaims {
	return jwt.MapClaims{
		"userId":   userId,
		"userName": username,
		"email":    email,
		"exp":      expirationTime,
	}
}

func VerifyJWT(request *http.Request) (bool, error) {
	_, err := request.Cookie("ghs_token")

	if err == http.ErrNoCookie {
		return false, err
	}

	return true, nil
}
