package cfg

import (
	"time"

	"github.com/Deirror/dutils/env"
)

// DBConfig holds the configuration parameters for a database connection.
type DBConfig struct {
	DSN         string        // DSN is used to connect to the database
	PoolSize    uint8         // Maximum number of open connections in the pool
	MaxIdle     uint8         // Maximum number of idle connections in the pool
	MaxLifetime time.Duration // Maximum lifetime of a connection before it's recycled
}

func NewDBConfig(dsn string, size, maxIdle uint8, maxLT time.Duration) *DBConfig {
	return &DBConfig{
		DSN:         dsn,
		PoolSize:    size,
		MaxIdle:     maxIdle,
		MaxLifetime: maxLT,
	}
}

// LoadEnvDBConfig loads the database configuration from environment variables.
// Required variables: DB_DSN, DB_POOL_SIZE, DB_MAX_IDLE, DB_MAX_LIFETIME.
func LoadEnvDBConfig() (*DBConfig, error) {
	dsn, err := env.GetEnv("DB_DSN")
	if err != nil {
		return nil, err
	}

	size, err := env.ParseEnvInt("DB_POOL_SIZE")
	if err != nil {
		return nil, err
	}

	maxIdle, err := env.ParseEnvInt("DB_MAX_IDLE")
	if err != nil {
		return nil, err
	}

	maxLT, err := env.ParseEnvTimeDuration("DB_MAX_LIFETIME")
	if err != nil {
		return nil, err
	}

	return NewDBConfig(dsn, uint8(size), uint8(maxIdle), maxLT*time.Second), nil
}

// WithPoolSize sets the PoolSize and returns the updated DBConfig.
func (cfg *DBConfig) WithPoolSize(size uint8) *DBConfig {
	cfg.PoolSize = size
	return cfg
}

// WithIdle sets the MaxIdle value and returns the updated DBConfig.
func (cfg *DBConfig) WithIdle(idle uint8) *DBConfig {
	cfg.MaxIdle = idle
	return cfg
}

// WithMaxLifetime sets the MaxLifetime value and returns the updated DBConfig.
func (cfg *DBConfig) WithMaxLifetime(maxTL time.Duration) *DBConfig {
	cfg.MaxLifetime = maxTL
	return cfg
}

// WithDSN sets the DSN string and returns the updated DBConfig.
func (cfg *DBConfig) WithDSN(dsn string) *DBConfig {
	cfg.DSN = dsn
	return cfg
}
