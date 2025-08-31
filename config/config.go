package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DB         DBConfig
	Redis      RedisConfig
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

type RedisConfig struct {
	Addr     string
	Port     string
	Password string
	DB       int
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
		Redis: RedisConfig{
			Addr:     getEnv("REDIS_ADDR", "localhost"),
			Port:     getEnv("REDIS_Port", "6379"),
			Password: getEnv("REDIS_PASSWORD", ""),
			DB:       getEnv("REDIS_DB", 0),
		},
		MachineId:  getEnv("MACHINE_ID", "1"),
		ServerPort: getEnv("SERVER_PORT", "8989"),
		JWTSecret:  getEnv("JWT_SECRET", "XgpHzFtngBNuyCnGXVZYlf5znjkXWsDs"),
	}, nil
}

type EnvType interface {
	string | int | float64 | bool
}

func getEnv[T EnvType](key string, defaultValue T) T {
	if value, exists := os.LookupEnv(key); exists {
		switch any(defaultValue).(type) {
		case int:
			i, err := strconv.Atoi(value)
			if err != nil {
				return defaultValue
			}
			return any(i).(T)
		case string:
			return any(value).(T)
		case bool:
			b, err := strconv.ParseBool(value)
			if err != nil {
				return defaultValue
			}
			return any(b).(T)
		case float64:
			f, err := strconv.ParseFloat(value, 64)
			if err != nil {
				return defaultValue
			}
			return any(f).(T)
		default:
			return defaultValue
		}
	}
	return defaultValue
}
