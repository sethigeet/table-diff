package database

import (
	"database/sql"
	"fmt"
	"os"
)

func Connect() error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_DBNAME"),
		os.Getenv("DB_PORT"),
	)

	var err error
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		return err
	}

	return nil
}
