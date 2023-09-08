package handlers

import (
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

	router.HandleFunc("/house/{id}", findHouseHandler.FindHouseHandler).Methods("GET")
	router.HandleFunc("/house", createHouseHandler.CreateHouseHandler).Methods("POST")
	router.HandleFunc("/house/{id}", deleteHouseHandler.DeleteHouseHandler).Methods("DELETE")

}
