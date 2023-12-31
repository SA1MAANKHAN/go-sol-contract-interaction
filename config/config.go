package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type ConfigType struct {
	PRIVATE_KEY  string
	PROVIDER_RPC string
}

var Config *ConfigType

func LoadConfig() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Config = &ConfigType{
		PRIVATE_KEY:  string(os.Getenv("PRIVATE_KEY")),
		PROVIDER_RPC: string(os.Getenv("PROVIDER_RPC")),
	}
}
