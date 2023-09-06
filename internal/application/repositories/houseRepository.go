package repositories

import (
	"github.com/luksrocha/house-system/internal/domain"
)

type HouseRepository interface {
	Insert(house *domain.House) error
	Delete(id string) error
	Find(id string) (*domain.House, error)
	Update(house *domain.House) error
}
