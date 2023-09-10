package useCase

import (
	"github.com/luksrocha/house-system/internal/domain/entities"
	"github.com/luksrocha/house-system/internal/domain/repositories"
	"github.com/luksrocha/house-system/internal/infra/dto"
)

type UpdateHouseUseCase struct {
	houseRepository repositories.HouseRepository
}

func NewUpdateHouseUseCase(repo repositories.HouseRepository) *UpdateHouseUseCase {
	return &UpdateHouseUseCase{
		houseRepository: repo,
	}
}

func (u *UpdateHouseUseCase) Execute(id string, data dto.CreateHouseDTOInput) (*entities.House, error) {
	houseFound, err := u.houseRepository.Find(id)

	if err != nil {
		return nil, err
	}

	houseFound.Name = data.Name
	houseFound.Address = data.Address

	houseUpdated, err := u.houseRepository.Update(houseFound)

	if err != nil {
		return nil, err
	}

	return houseUpdated, nil

}
