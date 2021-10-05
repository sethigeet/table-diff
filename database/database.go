// Package database provides the functionality to connect to the database
package database

import (
	"database/sql"
	"log"
)

var DB *sql.DB

func init() {
	if err := Connect(); err != nil {
		log.Fatalf("errors while connecting to the database: \n%s", err)
	}
}

type Row []interface{}

type Table []Row
