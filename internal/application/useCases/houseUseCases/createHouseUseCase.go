package useCase

import (
	"github.com/luksrocha/house-system/internal/domain/entities"
	"github.com/luksrocha/house-system/internal/domain/repositories"
)

type CreateHouseUseCase struct {
	houseRepository repositories.HouseRepository
}

func NewCreateHouseUseCase(houseRepository repositories.HouseRepository) *CreateHouseUseCase {
	return &CreateHouseUseCase{
		houseRepository: houseRepository,
	}
}

func (useCase *CreateHouseUseCase) Execute(house *entities.House) error {

	err := useCase.houseRepository.Insert(house)

	if err != nil {
		return err
	}

	return nil

}
