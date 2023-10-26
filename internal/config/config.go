package config

import (
	"os"
)

type Config struct {
	Port            string
	MongoURL        string
	Database        string
	UserCollection  string
	PasteCollection string
}

func NewConfig() *Config {
	return &Config{
		Port:            getEnv("PORT", ":8080"),
		MongoURL:        getEnv("MONGO_URL", ""),
		Database:        getEnv("DATABASE", ""),
		UserCollection:  getEnv("USER_COLLECTION", ""),
		PasteCollection: getEnv("PASTE_COLLECTION", ""),
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
