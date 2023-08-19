package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// It loads the .env file and returns the value of the key passed to it
func Config(key string) (string, error) {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Print("Error loading .env file ", err)
		return "", err
	}
	fmt.Println("Cargando .env")
	return os.Getenv(key), nil
}
