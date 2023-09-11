package userHandlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/luksrocha/house-system/internal/domain/repositories"
)

type UserHandler struct {
	repo repositories.UserRepository
}

func NewUserHandler(repo repositories.UserRepository) *UserHandler {
	return &UserHandler{repo: repo}
}

func (u *UserHandler) RegisterHandler(router *mux.Router) {
	createUserHandler := NewCreateUserHandler(u.repo)
	loginUserHandler := NewLoginUserHandler(u.repo)

	r := router.PathPrefix("/user").Subrouter()

	r.HandleFunc("", createUserHandler.Execute).Methods(http.MethodPost)
	r.HandleFunc("/login", loginUserHandler.Execute).Methods(http.MethodPost)

}
