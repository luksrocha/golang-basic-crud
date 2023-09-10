package userHandlers

import (
	"encoding/json"
	"net/http"

	useCase "github.com/luksrocha/house-system/internal/application/useCases/userUseCases"
	"github.com/luksrocha/house-system/internal/domain/entities"
	"github.com/luksrocha/house-system/internal/domain/repositories"
	"github.com/luksrocha/house-system/internal/infra/dto"
)

type CreateUserHandler struct {
	repo repositories.UserRepository
}

func NewCreateUserHandler(repo repositories.UserRepository) *CreateUserHandler {
	return &CreateUserHandler{
		repo: repo,
	}
}

func (c *CreateUserHandler) Execute(response http.ResponseWriter, request *http.Request) {
	var user dto.CreateUserDTOInput

	err := json.NewDecoder(request.Body).Decode(&user)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	createUserUseCase := useCase.NewCreateUserUseCase(c.repo)

	entityUser := entities.NewUser(user.FirstName, user.LastName, user.Email, user.Password)

	err = createUserUseCase.Execute(entityUser)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	userJson, err := json.Marshal(entityUser)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusCreated)
	response.Write(userJson)

}
