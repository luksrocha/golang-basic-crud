package main

import (
	"net/http"

	_ "github.com/lib/pq"
	"github.com/luksrocha/house-system/database"
	"github.com/luksrocha/house-system/internal/application/repositories"
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

	http.HandleFunc("/createHouse", createHouseHandler.CreateHouseHandler())
	http.HandleFunc("/deleteHouse", deleteHouseHandler.DeleteHouseHandler())

	http.ListenAndServe(":8090", nil)

}
