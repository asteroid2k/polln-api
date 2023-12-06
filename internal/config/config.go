package config

import (
	"os"
	"strconv"
)

type App struct {
	Env *Env
}

type DBConfig struct {
}

type Env struct {
	ReleaseStage string
	Port         int
	DB           DBConfig
}

func ParseEnv() *Env {
	return &Env{
		ReleaseStage: getEnv("ENV", "development"),
		Port:         getEnv("PORT", 4069),
		DB:           DBConfig{},
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
