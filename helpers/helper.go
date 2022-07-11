package helpers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetDsn(key string) string {

	err := godotenv.Load(".env")
  
	if err != nil {
	  log.Fatal("Error loading .env file")
	}
  
	return os.Getenv(key)
}