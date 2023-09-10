package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	useCase "github.com/luksrocha/house-system/internal/application/useCases/houseUseCases"
	"github.com/luksrocha/house-system/internal/domain/repositories"
	"github.com/luksrocha/house-system/internal/infra/dto"
)

type UpdateHouseHandler struct {
	repo repositories.HouseRepository
}

func NewUpdateHouseHandler(repo repositories.HouseRepository) *UpdateHouseHandler {
	return &UpdateHouseHandler{
		repo: repo,
	}
}

func (h *UpdateHouseHandler) Handle(response http.ResponseWriter, request *http.Request) {
	updateHouseUseCase := useCase.NewUpdateHouseUseCase(h.repo)

	vars := mux.Vars(request)

	id, ok := vars["id"]

	if !ok {
		response.Write([]byte("Missing id parameter"))
		return
	}

	var house dto.CreateHouseDTOInput

	err := json.NewDecoder(request.Body).Decode(&house)

	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	updateHouseUseCase.Execute(id, house)
}
