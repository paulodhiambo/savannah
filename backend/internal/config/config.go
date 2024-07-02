package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type Config struct {
	Port               string
	DBHost             string
	DBPort             int
	DBUser             string
	DBPassword         string
	DBName             string
	SMSSandboxAPIKey   string
	SMSSandboxUserName string
	GithubClientID     string
	GithubClientSecret string
	CallbackUrl        string
}

var AppConfig Config

func Load() error {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables instead")
	}

	dbPort, err := strconv.Atoi(getEnv("DB_PORT", "5432"))
	if err != nil {
		log.Fatalf("Invalid port number for DB_PORT: %v", err)
	}

	AppConfig = Config{
		Port:               getEnv("PORT", "8080"),
		DBHost:             getEnv("DB_HOST", "localhost"),
		DBPort:             dbPort,
		DBUser:             getEnv("DB_USER", "user"),
		DBPassword:         getEnv("DB_PASSWORD", "password"),
		DBName:             getEnv("DB_NAME", "database"),
		SMSSandboxAPIKey:   getEnv("SMS_SANDBOX_API_KEY", ""),
		SMSSandboxUserName: getEnv("SMS_SANDBOX_API_USERNAME", ""),
		GithubClientID:     getEnv("CLIENT_ID", ""),
		GithubClientSecret: getEnv("CLIENT_SECRET", ""),
		CallbackUrl:        getEnv("CALL_BACK_URL", ""),
	}

	log.Printf("Configuration loaded successfully %+v", AppConfig)
	return nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
