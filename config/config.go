package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Host     string
	Port     string
	Password string
	User     string
	DBName   string
	SSLMode  string
	Link     string
	HostPort string
	JWT_SIGN_KEY string
}

func LoadConfig() (Config, error) {
	var config Config

	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, using default environment variables")
	}

	// Get environment variables
	config.Host = getEnv("HOST", "localhost")         // default to localhost
	config.Port = getEnv("PORT", "5432")              // default to 5432
	config.User = getEnv("USER", "default")           // default to "default"
	config.Password = getEnv("PASSWORD", "")          // default to empty
	config.DBName = getEnv("DBName", "edumarshal_db") // default to "edumarshal_db"
	config.SSLMode = getEnv("SSLMode", "require")     // default to "disable"
	config.Link = getEnv("LINK", "http://localhost")  // default to "http://localhost"
	config.HostPort = getEnv("HOSTPORT", "3000") // default
	config.JWT_SIGN_KEY = getEnv("JWT_SIGN_KEY", "") // default to ""
	return config, nil
}

// getEnv retrieves the value of the environment variable key or returns the fallback value if it doesn't exist.
func getEnv(key string, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}
	return value
}
