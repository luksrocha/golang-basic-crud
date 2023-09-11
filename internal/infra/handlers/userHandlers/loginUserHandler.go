package userHandlers

import (
	"encoding/json"
	"net/http"
	"time"

	useCase "github.com/luksrocha/house-system/internal/application/useCases/authenticateUseCase"
	"github.com/luksrocha/house-system/internal/domain/repositories"
	"github.com/luksrocha/house-system/internal/infra/dto"
)

type LoginUserHandler struct {
	repo repositories.UserRepository
}

func NewLoginUserHandler(repo repositories.UserRepository) *LoginUserHandler {
	return &LoginUserHandler{
		repo: repo,
	}
}

func (l *LoginUserHandler) Execute(response http.ResponseWriter, request *http.Request) {
	var loginInput dto.LoginDTOInput

	err := json.NewDecoder(request.Body).Decode(&loginInput)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	loginUseCase := useCase.NewLoginUserUseCase(l.repo)

	token, err := loginUseCase.Execute(loginInput.Email, loginInput.Password)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	cookie := http.Cookie{}
	cookie.Name = "ghs_token"
	cookie.Value = token
	cookie.Expires = time.Now().Add(365 * 24 * time.Hour)
	cookie.Secure = false
	cookie.HttpOnly = true
	cookie.Path = "/"

	response.Header().Set("Content-Type", "application/json")
	http.SetCookie(response, &cookie)
	response.WriteHeader(http.StatusOK)
	response.Write([]byte(token))
}
