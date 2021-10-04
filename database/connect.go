package database

import (
	"database/sql"
	"fmt"

	// Load the postgres driver for database/sql
	_ "github.com/lib/pq"

	"github.com/sethigeet/table-diff/config"
)

func Connect() error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		config.Config.DB.Host,
		config.Config.DB.User,
		config.Config.DB.Password,
		config.Config.DB.DBName,
		config.Config.DB.Port,
		config.Config.DB.SSLMode,
	)

	var err error
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		return err
	}

	return nil
}
