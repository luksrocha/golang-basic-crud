package main

import (
	"fmt"
	"net/http"
	"time"

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

	router := mux.NewRouter().StrictSlash(true)

	houseRepository := repositories.NewHouseRepositoryPostgres(db)

	houseHandler := handlers.NewHouseHandler(&houseRepository)
	houseHandler.RegisterHandlers(router)

	srv := &http.Server{
		Addr: ":8090",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router, // Pass our instance of gorilla/mux in.
	}

	if err := srv.ListenAndServe(); err != nil {
		fmt.Println(err)
	}

}
