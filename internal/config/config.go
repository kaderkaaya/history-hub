package config

import (
	"os"
	"path/filepath"
	"runtime"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort          string
	AppEnv           string
	RedisPort        string
	RedisPassword    string
	RedisHost        string
	WikimediaBaseURL string
	CacheTTLTodayH   int
	CacheTTLPastH    int
}

func Load() Config {
	_, b, _, _ := runtime.Caller(0)
	rootDir := filepath.Join(filepath.Dir(b), "..", "..")
	_ = godotenv.Load(filepath.Join(rootDir, ".env"))
	return Config{
		AppPort:          getEnv("APP_PORT", "8080"),
		AppEnv:           getEnv("APP_ENV", "development"),
		RedisPort:        getEnv("REDIS_PORT", ""),
		RedisPassword:    getEnv("REDIS_PASSWORD", ""),
		RedisHost:        getEnv("REDIS_HOST", ""),
		WikimediaBaseURL: getEnv("WIKIMEDIA_BASE_URL", "https://api.wikimedia.org"),
		CacheTTLTodayH:   getEnvInt("CACHE_TTL_TODAY_HOURS", 12),
		CacheTTLPastH:    getEnvInt("CACHE_TTL_PAST_HOURS", 168),
	}
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

func getEnvInt(key string, fallback int) int {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	n, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}
	return n
}
