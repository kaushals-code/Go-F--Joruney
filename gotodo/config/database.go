package config

import (
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func ConnectDB() (*sql.DB, error) {
	connString := "postgres://postgres:12345@localhost:5433/practice"

	db, err := sql.Open("pgx", connString)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil

}
