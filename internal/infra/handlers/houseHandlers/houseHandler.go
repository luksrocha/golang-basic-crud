package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/luksrocha/house-system/internal/domain/repositories"
)

type HouseHandler struct {
	repo repositories.HouseRepository
}

func NewHouseHandler(repo repositories.HouseRepository) *HouseHandler {
	return &HouseHandler{
		repo: repo,
	}
}

func (h *HouseHandler) RegisterHandlers(router *mux.Router) {
	createHouseHandler := NewCreateHouseHandler(h.repo)
	deleteHouseHandler := NewDeleteHouseHandler(h.repo)
	findHouseHandler := NewFindHouseHandler(h.repo)
	updateHouseHandler := NewUpdateHouseHandler(h.repo)

	r := router.PathPrefix("/house").Subrouter()

	r.HandleFunc("", createHouseHandler.CreateHouseHandler).Methods(http.MethodPost)
	r.HandleFunc("/{id}", deleteHouseHandler.DeleteHouseHandler).Methods(http.MethodDelete)
	r.HandleFunc("/{id}", findHouseHandler.FindHouseHandler).Methods(http.MethodGet)
	r.HandleFunc("/{id}", updateHouseHandler.Handle).Methods(http.MethodPut)

}
