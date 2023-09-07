package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func OpenConnection() (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", "host=localhost user=postgres password=bob dbname=ghs sslmode=disable")

	if err != nil {
		panic(err)
	}

	return db, err
}
