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

	AccessTokenPrivateKey  string
	AccessTokenPublicKey   string
	RefreshTokenPrivateKey string
	RefreshTokenPublicKey  string
}

func NewConfig() *Config {
	return &Config{
		Port:            getEnv("PORT", ":8080"),
		MongoURL:        getEnv("MONGO_URL", "mongodb://localhost:27017/"),
		Database:        getEnv("DATABASE", "golang"),
		UserCollection:  getEnv("USER_COLLECTION", "users"),
		PasteCollection: getEnv("PASTE_COLLECTION", "pastes"),

		AccessTokenPrivateKey:  getEnv("ACCESS_TOKEN_PRIVATE_KEY", ""),
		AccessTokenPublicKey:   getEnv("ACCESS_TOKEN_PUBLIC_KEY", ""),
		RefreshTokenPrivateKey: getEnv("REFRESH_TOKEN_PRIVATE_KEY", ""),
		RefreshTokenPublicKey:  getEnv("REFRESH_TOKEN_PUBLIC_KEY", ""),
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
