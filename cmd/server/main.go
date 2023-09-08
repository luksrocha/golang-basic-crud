package main

import (
	"net/http"

	"github.com/gorilla/mux"
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

	mux := mux.NewRouter()

	houseRepository := repositories.NewHouseRepositoryPostgres(db)

	houseHandler := handlers.NewHouseHandler(&houseRepository)

	houseHandler.RegisterHandlers(mux)

	srv := http.NewServeMux()

	srv.Handle("/", mux)

	http.ListenAndServe(":8090", mux)

}
