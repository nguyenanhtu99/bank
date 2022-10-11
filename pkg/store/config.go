package store

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoUrl		string
	MongoDatabase	string
}

func loadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}
	cfg := Config{
		MongoUrl: os.Getenv("MONGO_URL"),
		MongoDatabase: os.Getenv("MONGO_DATABASE"),
	}
	return &cfg, nil
}