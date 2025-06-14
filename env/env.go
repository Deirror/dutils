package env

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

const (
	DEV     = "dev"
	PROD    = "prod"
	STAGING = "staging"
)

// Loads environment variables from the given files
// It is intented to be used in development mode
func LoadEnv(filenames ...string) error {
	if err := godotenv.Load(filenames...); err != nil {
		return fmt.Errorf("error loading .env file: %v", err)
	}
	return nil
}

// A wrapper func around os.Getenv, but handles error with more text
func GetEnv(key string) (string, error) {
	val := os.Getenv(key)

	if val == "" {
		return "", fmt.Errorf("environment variable %s not set", key)
	}

	return val, nil
}

// Environment variable MODE must be set in .env to call the func
// Can be used for dev, prod and staging environments
func GetMode() (string, error) {
	mode, err := GetEnv("MODE")
	if err != nil {
		return "", err
	}

	// Make it case-insensitive
	mode = strings.ToLower(mode)

	switch mode {
	case DEV:
	case PROD:
	case STAGING:
		return mode, nil
	}

	return "", fmt.Errorf("invalid environment mode: %s", mode)
}

// A wrapper func around os.Environ, handling errors more precisely
func GetAllEnv() ([]string, error) {
	kvps := os.Environ()
	if len(kvps) == 0 {
		return nil, fmt.Errorf("no environment variables found")
	}
	return kvps, nil
}
