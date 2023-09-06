package useCase

import (
	"github.com/luksrocha/house-system/internal/application/repositories"
	"github.com/luksrocha/house-system/internal/domain"
)

type CreateHouseUseCase struct {
	houseRepository repositories.HouseRepository
}

func NewCreateHouseUseCase(houseRepository repositories.HouseRepository) *CreateHouseUseCase {
	return &CreateHouseUseCase{
		houseRepository: houseRepository,
	}
}

func (useCase *CreateHouseUseCase) Execute(house *domain.House) error {

	err := useCase.houseRepository.Insert(house)

	if err != nil {
		return err
	}

	return nil

}
