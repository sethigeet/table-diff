// Package config provides the necessary functions to create a configuration
// object for the application and properly load config values
package config

import (
	"log"
	"os"
	"strconv"
)

type DBConfigType struct {
	SSLMode, Host, User, Password, DBName string
	Port                                  uint16
	Table1Name, Table2Name                string
}

type DiffAlgorithmConfigType struct {
	KeyColumns    []string
	IgnoreColumns []string
}

type ConfigType struct {
	DB            DBConfigType
	DiffAlgorithm DiffAlgorithmConfigType
}

var Config ConfigType

func init() {
	// Load environment variables
	if err := LoadEnv(true); err != nil {
		log.Fatalf("there were errors while loading the env file: \n%s", err)
	}

	port, err := strconv.ParseUint(os.Getenv("DB_PORT"), 10, 16)
	if err != nil {
		log.Fatalf("there were errors while loading the env file: \n%s", err)
	}

	Config = ConfigType{
		DB: DBConfigType{
			SSLMode:  os.Getenv("DB_SSL_MODE"),
			Host:     os.Getenv("DB_HOST"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			DBName:   os.Getenv("DB_DBNAME"),
			Port:     uint16(port),

			// TODO: Don't hardcode this
			Table1Name: "mock_data_1",
			Table2Name: "mock_data_2",
		},
		DiffAlgorithm: DiffAlgorithmConfigType{
			// TODO: Don't hardcode this
			KeyColumns:    []string{"name", "username"},
			IgnoreColumns: []string{"id"},
		},
	}
}
