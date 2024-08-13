package initializers

import (
	"log"
	"path/filepath"

	"github.com/joho/godotenv"
)

func LoadENV() {

	envFilePath := filepath.Join(".env")
	if err := godotenv.Load(envFilePath); err != nil {
		log.Fatalf("Error loading %s file: %v", envFilePath, err)
	}

}
