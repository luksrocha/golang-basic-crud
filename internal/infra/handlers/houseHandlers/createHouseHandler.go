package handlers

import (
	"encoding/json"
	"net/http"

	useCase "github.com/luksrocha/house-system/internal/application/useCases/houseUseCases"
	"github.com/luksrocha/house-system/internal/domain/entities"
	"github.com/luksrocha/house-system/internal/domain/repositories"
	"github.com/luksrocha/house-system/internal/infra/dto"
)

type CreateHouseHandler struct {
	HouseDB repositories.HouseRepository
}

func NewCreateHouseHandler(db repositories.HouseRepository) *CreateHouseHandler {
	return &CreateHouseHandler{
		HouseDB: db,
	}
}

func (h *CreateHouseHandler) CreateHouseHandler(response http.ResponseWriter, request *http.Request) {

	var house dto.CreateHouseDTOInput

	err := json.NewDecoder(request.Body).Decode(&house)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	createHouseUseCase := useCase.NewCreateHouseUseCase(h.HouseDB)

	domainHouse, err := entities.NewHouse(house.Name, house.Address)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = createHouseUseCase.Execute(domainHouse)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	houseJson, _ := json.Marshal(domainHouse)
	response.Write(houseJson)

}
