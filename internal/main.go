package main

import (
	"net/http"

	_ "github.com/lib/pq"
	"github.com/luksrocha/house-system/database"
	"github.com/luksrocha/house-system/internal/application/repositories"
	"github.com/luksrocha/house-system/internal/infra/handlers"
)

func main() {

	db, err := database.OpenConnection()

	if err != nil {
		panic(err)
	}

	defer db.Close()

	houseRepository := repositories.NewHouseRepositoryPostgres(db)

	createHouseHandler := handlers.NewHouseHandler(&houseRepository)

	http.HandleFunc("/teste", createHouseHandler.CreateHouseHandler())

	http.ListenAndServe(":8090", nil)

}
