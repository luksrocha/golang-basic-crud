package useCase

import (
	"errors"

	"github.com/luksrocha/house-system/internal/domain/repositories"
	"golang.org/x/crypto/bcrypt"
)

type LoginUserUseCase struct {
	repo repositories.UserRepository
}

func NewLoginUserUseCase(repo repositories.UserRepository) *LoginUserUseCase {
	return &LoginUserUseCase{repo: repo}
}

func (l *LoginUserUseCase) Execute(email, password string) (string, error) {
	user, err := l.repo.FindByEmail(email)

	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password))

	if err != nil {
		return "", errors.New("invalid password or email, please, try again")
	}

	token, err := user.GenerateToken()

	if err != nil {
		return "", err
	}

	return token, nil

}
