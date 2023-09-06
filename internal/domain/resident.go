package domain

import (
	"github.com/asaskevich/govalidator"
)

type Resident struct {
	ID    string `json:"id" type:"uuid"`
	Name  string `json:"name" type:"string"`
	Email string `json:"email"`
	House *House
}

func (resident *Resident) Validate() error {
	_, err := govalidator.ValidateStruct(resident)

	if err != nil {
		return err
	}

	return nil
}

func NewResident() *Resident {
	resident := Resident{}
	return &resident
}

func HasValidHouse(resident Resident) bool {
	return resident.House != nil
}
