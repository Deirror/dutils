package env

import (
	"fmt"
	"os"
	"strconv"
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
// If in dev mode, make sure to load the env vars first, then call other funcs
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
		return mode, err
	}

	switch strings.ToLower(mode) {
	case DEV, PROD, STAGING:
		return mode, nil
	default:
		return mode, fmt.Errorf("invalid environment mode: %s", mode)
	}
}

// A wrapper func around os.Environ, handling errors more precisely
func GetAllEnv() ([]string, error) {
	kvps := os.Environ()
	if len(kvps) == 0 {
		return nil, fmt.Errorf("no environment variables found")
	}
	return kvps, nil
}

// Same as GetEnv, but with default value
func GetEnvOrDefault(key, defaultVal string) string {
	val := os.Getenv(key)
	if val == "" {
		return defaultVal
	}
	return val
}

// Gets key's value, calling GetEnv and parses to bool
// Assumes it is in a specific format - text or a binary representation
func ParseEnvBool(key string) (bool, error) {
	val, err := GetEnv(key)
	if err != nil {
		return false, err
	}

	switch strings.ToLower(val) {
	case "true", "1", "yes", "y":
		return true, nil
	case "false", "0", "no", "n":
		return false, nil
	default:
		return false, fmt.Errorf("invalid boolean value for %s: %s", key, val)
	}
}

// Gets and parses key's value to int
func ParseEnvInt(key string) (int, error) {
	val, err := GetEnv(key)
	if err != nil {
		return 0, err
	}

	i, err := strconv.Atoi(val)
	if err != nil {
		return 0, fmt.Errorf("invalid integer value for %s: %s", key, val)
	}
	return i, nil
}
