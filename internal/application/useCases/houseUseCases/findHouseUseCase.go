package useCase

import (
	"github.com/luksrocha/house-system/internal/domain/entities"
	"github.com/luksrocha/house-system/internal/domain/repositories"
)

type FindHouseUseCase struct {
	houseRepository repositories.HouseRepository
}

func NewFindHouseUseCase(repo repositories.HouseRepository) *FindHouseUseCase {
	return &FindHouseUseCase{
		houseRepository: repo,
	}
}

func (f *FindHouseUseCase) Exec(id string) (*entities.House, error) {
	house, err := f.houseRepository.Find(id)

	if err != nil {
		return nil, err
	}

	return house, nil
}
