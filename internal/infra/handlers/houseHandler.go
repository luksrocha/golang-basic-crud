package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/luksrocha/house-system/internal/application/repositories"
	useCase "github.com/luksrocha/house-system/internal/application/useCases/houseUseCases"
	"github.com/luksrocha/house-system/internal/domain"
	"github.com/luksrocha/house-system/internal/infra/dto"
)

type HouseHandler struct {
	HouseDB repositories.HouseRepository
}

func NewHouseHandler(db repositories.HouseRepository) *HouseHandler {
	return &HouseHandler{
		HouseDB: db,
	}
}

func (h *HouseHandler) CreateHouseHandler() func(response http.ResponseWriter, request *http.Request) {
	return func(response http.ResponseWriter, request *http.Request) {

		var house dto.CreateHouseDTOInput

		err := json.NewDecoder(request.Body).Decode(&house)

		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			return
		}

		createHouseUseCase := useCase.NewCreateHouseUseCase(h.HouseDB)

		domainHouse, err := domain.NewHouse(house.Name, house.Address)

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

}
