package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB         DBConfig
	MachineId  string
	ServerPort string
	JWTSecret  string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func Load() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	return &Config{
		DB: DBConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "magic"),
			Password: getEnv("DB_PASSWORD", "magic@520"),
			DBName:   getEnv("DB_NAME", "ixpay"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		},
		MachineId:  getEnv("MACHINE_ID", "1"),
		ServerPort: getEnv("SERVER_PORT", "8989"),
		JWTSecret:  getEnv("JWT_SECRET", "XgpHzFtngBNuyCnGXVZYlf5znjkXWsDs"),
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
