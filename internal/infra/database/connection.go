package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func OpenConnection(driverName, host, user, password, dbName string) (*sqlx.DB, error) {
	dsnName := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, password, dbName)

	db, err := sqlx.Connect(driverName, dsnName)

	if err != nil {
		panic(err)
	}

	return db, err
}
