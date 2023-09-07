package repositories

import (
	"github.com/luksrocha/house-system/internal/domain/entities"
)

type HouseRepository interface {
	Insert(house *entities.House) error
	Delete(id string) error
	Find(id string) (*entities.House, error)
	Update(house *entities.House) error
}
