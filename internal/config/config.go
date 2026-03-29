package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort          string
	AppEnv           string
	RedisPort        string
	RedisPassword    string
	RedisHost        string
	WikimediaBaseURL string
}

func Load() Config {
	_ = godotenv.Load() //eğer değiskeni bi daha kullanmayacaksak _ yap.
	return Config{
		AppPort:          getEnv("APP_PORT", "8080"),
		AppEnv:           getEnv("APP_ENV", "development"),
		RedisPort:        getEnv("REDIS_PORT", ""),
		RedisPassword:    getEnv("REDIS_PASSWORD", ""),
		RedisHost:        getEnv("REDIS_HOST", ""),
		WikimediaBaseURL: getEnv("WIKIMEDIA_BASE_URL", "https://api.wikimedia.org"),
	}
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
