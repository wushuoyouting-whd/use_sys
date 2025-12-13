package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
	ServerPort string
	BeType     string
}

var AppConfig *Config

func LoadConfig() {
	// 加载 .env 文件（如果存在）
	_ = godotenv.Load()

	port, _ := strconv.Atoi(getEnv("DB_PORT", "5432"))

	AppConfig = &Config{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     port,
		DBUser:     getEnv("DB_USER", "owen"),
		DBPassword: getEnv("DB_PASSWORD", "123456"),
		DBName:     getEnv("DB_NAME", "user_sys"),
		DBSSLMode:    getEnv("DB_SSLMODE", "disable"),
		ServerPort:   getEnv("SERVER_PORT", "3013"),
		BeType:       getEnv("BE_TYPE", "Go"),
	}

	log.Println("配置加载成功")
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

