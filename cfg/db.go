package cfg

import (
	"time"

	"github.com/Deirror/dutils/env"
)

type DBConfig struct {
	DSN         string
	PoolSize    uint8
	MaxIdle     uint8
	MaxLifetime time.Duration
}

func NewDBConfig(dsn string, size, maxIdle uint8, maxLT time.Duration) *DBConfig {
	return &DBConfig{
		DSN:         dsn,
		PoolSize:    size,
		MaxIdle:     maxIdle,
		MaxLifetime: maxLT,
	}
}

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

func (cfg *DBConfig) WithPoolSize(size uint8) *DBConfig {
	cfg.PoolSize = size
	return cfg
}

func (cfg *DBConfig) WithIdle(idle uint8) *DBConfig {
	cfg.MaxIdle = idle
	return cfg
}

func (cfg *DBConfig) WithMaxLifetime(maxTL time.Duration) *DBConfig {
	cfg.MaxLifetime = maxTL
	return cfg
}

func (cfg *DBConfig) WithDSN(dsn string) *DBConfig {
	cfg.DSN = dsn
	return cfg
}
