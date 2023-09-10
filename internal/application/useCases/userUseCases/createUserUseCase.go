package useCase

import (
	"errors"

	"github.com/luksrocha/house-system/internal/domain/entities"
	"github.com/luksrocha/house-system/internal/domain/repositories"
)

type CreateUserUseCase struct {
	repo repositories.UserRepository
}

func NewCreateUserUseCase(repo repositories.UserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{
		repo: repo,
	}
}

func (c *CreateUserUseCase) Execute(user *entities.User) error {
	_, err := c.repo.FindByEmail(user.Email)

	if err == nil {
		return errors.New("user already exists")
	}

	err = c.repo.Insert(user)

	if err != nil {
		return err
	}

	return nil
}
