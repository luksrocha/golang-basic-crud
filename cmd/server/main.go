package main

import (
	"net/http"

	"github.com/luksrocha/house-system/internal/domain/repositories"
	"github.com/luksrocha/house-system/internal/infra/database"
	handlers "github.com/luksrocha/house-system/internal/infra/handlers/houseHandlers"
)

func main() {

	db, err := database.OpenConnection()

	if err != nil {
		panic(err)
	}

	defer db.Close()

	houseRepository := repositories.NewHouseRepositoryPostgres(db)

	createHouseHandler := handlers.NewHouseHandler(&houseRepository)
	deleteHouseHandler := handlers.NewDeleteHouseHandler(&houseRepository)
	findHouseHandler := handlers.NewFindHouseHandler(&houseRepository)

	http.HandleFunc("/createHouse", createHouseHandler.CreateHouseHandler)
	http.HandleFunc("/deleteHouse", deleteHouseHandler.DeleteHouseHandler)
	http.HandleFunc("/findHouse", findHouseHandler.FindHouseHandler)

	http.ListenAndServe(":8090", nil)

}
