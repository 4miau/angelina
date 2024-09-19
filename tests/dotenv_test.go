package tests

import (
	"fmt"
	"github.com/joho/godotenv"
	"testing"
)

// Ensures dotenv is able to be loaded for project from .env file
func TestDotenvLoad(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		t.Fatal("Error loading .env file")
	}
}

func TestDotenvMap(t *testing.T) {
	envs, err := godotenv.Read("../.env")

	if err != nil {
		t.Fatal("Error loading .env file")
	} else {
		fmt.Printf("%+v\n", envs)
	}
}
