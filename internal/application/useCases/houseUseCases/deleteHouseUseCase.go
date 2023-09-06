package useCase

import "github.com/luksrocha/house-system/internal/application/repositories"

type DeleteHouseUseCase struct {
	houseRepository repositories.HouseRepository
}

func NewDeleteHouseUseCase(houseRepository repositories.HouseRepository) *DeleteHouseUseCase {
	return &DeleteHouseUseCase{
		houseRepository: houseRepository,
	}
}

func (d *DeleteHouseUseCase) Exec(id string) error {
	err := d.houseRepository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
