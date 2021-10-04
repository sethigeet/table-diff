package util

import (
	"fmt"

	"github.com/joho/godotenv"
)

// LoadEnv Loads the environment variables defined in the .env file according to
// the environment running
func LoadEnv(verify bool) error {
	envFile := ".env.local"

	if verify {
		err := verifyEnv(envFile)
		if err != nil {
			return err
		}
	}

	err := godotenv.Load(envFile)
	if err != nil {
		return err
	}

	return nil
}

// verifyEnv verifies that all the environment variables specified .env.example file
// are actually present in the environent and returns an error otherwise
func verifyEnv(filename string) error {
	var exampleEnv, actualEnv map[string]string
	var err error

	actualEnv, err = godotenv.Read(filename)
	if err != nil {
		return fmt.Errorf("an error occured while reading the '%s' file:\n%s", filename, err)
	}

	exampleEnv, err = godotenv.Read(".env.example")
	if err != nil {
		return fmt.Errorf("an error occured while reading the '.env.example' file:\n%s", err)
	}

	for envVar := range exampleEnv {
		if actualEnv[envVar] == "" {
			return fmt.Errorf("the env var %s is not present but is defined in the .env.example file", envVar)
		}
	}

	return nil
}
