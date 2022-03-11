package config

import (
	"os"

	"github.com/joho/godotenv"
)

var ErrLoadingEnvs error

func init() {
	ErrLoadingEnvs = godotenv.Load()
}

func init() {
}

func GetEnv(name, defVal string) string {
	env := os.Getenv(name)
	if env == "" {
		return defVal
	}

	return env
}
