package useCase

import (
	"errors"

	"github.com/luksrocha/house-system/internal/domain/repositories"
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

	if password != user.HashedPassword {
		return "", errors.New("invalid password or email, please, try again")
	}

	token, err := user.GenerateToken()

	if err != nil {
		return "", err
	}

	return token, nil

}
