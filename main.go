package main

import (
	"log"

	// Load the postgres driver for database/sql
	_ "github.com/lib/pq"

	"github.com/sethigeet/table-diff/database"
	"github.com/sethigeet/table-diff/util"
)

func main() {
	// Load environment variables
	if err := util.LoadEnv(true); err != nil {
		log.Fatalf("there were errors while loading the env file: \n%s", err)
	}

	err := database.Connect()
	if err != nil {
		log.Fatalf("errors while connecting to the database: \n%s", err)
	}
}
