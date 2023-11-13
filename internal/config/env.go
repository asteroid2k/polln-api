package config

import (
	"os"
	"strconv"
)

type DBConfig struct {
}

type Env struct {
	ReleaseStage string
	Port         int
	// DB  DBConfig
}

// New returns a new Config struct
func InitConfig() *Env {
	return &Env{
		ReleaseStage: getEnv("ENV", "development"),
		Port:         getEnv("PORT", 4069),
	}
}

func getEnv[T int | string](key string, defaultVal T) T {
	if valueStr, exists := os.LookupEnv(key); exists {

		switch any(defaultVal).(type) {
		case int:
			if value, err := strconv.Atoi(valueStr); err == nil {
				return T(rune(value))
			}
		}
	}
	return defaultVal
}
