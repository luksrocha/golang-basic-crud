package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/luksrocha/house-system/internal/domain"
)

type HouseRepositoryPostgres struct {
	DB *sqlx.DB
}

func NewHouseRepositoryPostgres(db *sqlx.DB) HouseRepositoryPostgres {
	return HouseRepositoryPostgres{DB: db}
}

func (houseRepository *HouseRepositoryPostgres) Insert(house *domain.House) error {
	prepare, err := houseRepository.DB.Prepare("INSERT INTO houses (id, name, address, created_at, updated_at ) VALUES ($1, $2, $3, $4, $5)")

	if err != nil {
		return err
	}

	_, err = prepare.Exec(house.ID, house.Name, house.Address, house.CreatedAt, house.UpdatedAt)

	if err != nil {
		return err
	}

	return nil

}

func (houseRepository *HouseRepositoryPostgres) Delete(id string) error {
	prepare, err := houseRepository.DB.Prepare("DELETE FROM houses WHERE id = $1")

	if err != nil {
		return err
	}

	_, err = prepare.Exec(id)

	if err != nil {
		return err
	}

	return nil
}

func (houseRepository *HouseRepositoryPostgres) Find(id string) (*domain.House, error) {
	prepare, err := houseRepository.DB.Prepare("SELECT * FROM houses WHERE id = $1")

	if err != nil {
		return nil, err
	}

	var house domain.House

	err = prepare.QueryRow(id).Scan(&house.ID, &house.Name, &house.Address, &house.CreatedAt, &house.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &house, nil

}

func (HouseRepository *HouseRepositoryPostgres) Update(house *domain.House) error {
	prepare, err := HouseRepository.DB.Prepare("UPDATE houses SET name = $1, address = $2, updated_at = $3 WHERE id = $4")

	if err != nil {
		return err
	}

	_, err = prepare.Exec(house.Name, house.Address, house.UpdatedAt, house.ID)

	if err != nil {
		return err
	}

	return nil
}
