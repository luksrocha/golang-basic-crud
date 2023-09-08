package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	useCase "github.com/luksrocha/house-system/internal/application/useCases/houseUseCases"
	"github.com/luksrocha/house-system/internal/domain/repositories"
)

type FindHouseHandler struct {
	repo repositories.HouseRepository
}

func NewFindHouseHandler(repo repositories.HouseRepository) *FindHouseHandler {
	return &FindHouseHandler{
		repo: repo,
	}
}

func (f *FindHouseHandler) FindHouseHandler(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	id, ok := vars["id"]

	if !ok {
		response.Write([]byte("Missing id parameter"))
		return
	}

	findHouseUseCase := useCase.NewFindHouseUseCase(f.repo)

	house, err := findHouseUseCase.Exec(id)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsonHouse, err := json.Marshal(house)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusOK)
	response.Write(jsonHouse)
}
