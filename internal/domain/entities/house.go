package entities

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/google/uuid"
)

type House struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (h *House) prepare() {
	h.ID = uuid.New()
	h.CreatedAt = time.Now()
	h.UpdatedAt = time.Now()
}

func (house *House) Validate() error {
	_, err := govalidator.ValidateStruct(house)

	if err != nil {
		return err
	}

	return nil
}

func NewHouse(name, address string) (*House, error) {

	house := House{}

	house.prepare()

	house.Name = name
	house.Address = address

	err := house.Validate()

	if err != nil {
		return nil, err
	}

	return &house, nil

}
