package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config 애플리케이션 설정 구조체
type Config struct {
	MongoURI      string
	MongoDatabase string
	ServerPort    string
	AppEnv        string
}

// LoadConfig 환경 변수를 로드하고 설정을 반환
func LoadConfig() *Config {
	// .env 파일 로드
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	}

	config := &Config{
		MongoURI:      getEnv("MONGO_URI", "mongodb://localhost:27017"),
		MongoDatabase: getEnv("MONGO_DATABASE", "blueberry_homework"),
		ServerPort:    getEnv("SERVER_PORT", "8080"),
		AppEnv:        getEnv("APP_ENV", "development"),
	}

	return config
}

// getEnv 환경 변수를 가져오고, 없으면 기본값을 반환
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
