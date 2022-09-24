package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetMongoURI() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error while loading .env file")
	}

	return os.Getenv("MONGOURI")
}
