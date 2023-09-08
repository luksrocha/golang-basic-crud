package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	useCase "github.com/luksrocha/house-system/internal/application/useCases/houseUseCases"
	"github.com/luksrocha/house-system/internal/domain/repositories"
)

type DeleteHouseHandler struct {
	repository repositories.HouseRepository
}

func NewDeleteHouseHandler(repository repositories.HouseRepository) *DeleteHouseHandler {
	return &DeleteHouseHandler{
		repository: repository,
	}
}

func (d *DeleteHouseHandler) DeleteHouseHandler(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	id, ok := vars["id"]

	if !ok {
		response.Write([]byte("Missing id parameter"))
		return
	}

	fmt.Println(id)

	if id == "" {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	deleteHouseUseCase := useCase.NewDeleteHouseUseCase(d.repository)

	if err := deleteHouseUseCase.Exec(id); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusOK)
	response.Write([]byte("House deleted successfully"))

}
