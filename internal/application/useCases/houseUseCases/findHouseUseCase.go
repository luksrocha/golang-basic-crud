package useCase

import (
	"github.com/luksrocha/house-system/internal/application/repositories"
	"github.com/luksrocha/house-system/internal/domain"
)

type FindHouseUseCase struct {
	houseRepository repositories.HouseRepository
}

func NewFindHouseUseCase(repo repositories.HouseRepository) *FindHouseUseCase {
	return &FindHouseUseCase{
		houseRepository: repo,
	}
}

func (f *FindHouseUseCase) Exec(id string) (*domain.House, error) {
	house, err := f.houseRepository.Find(id)

	if err != nil {
		return nil, err
	}

	return house, nil
}
