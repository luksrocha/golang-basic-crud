package repositories

import "github.com/luksrocha/house-system/internal/domain/entities"

type UserRepository interface {
	Insert(user *entities.User) error
	Delete(id string) error
	Find(id string) (*entities.User, error)
	Update(user *entities.User) (*entities.User, error)
	FindByEmail(email string) (*entities.User, error)
}
