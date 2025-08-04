package cfg

import (
	"time"

	"github.com/Deirror/dutils/env"
)

var dbSuffixes = []string{
	"DB_DRIVER",
	"DB_DSN",
	"DB_POOL_SIZE",
	"DB_MAX_IDLE",
	"DB_MAX_LIFETIME",
}

// DBConfig holds the configuration parameters for a database connection.
type DBConfig struct {
	Driver      string        // Driver is used to set db engine (postgres, mysql and so on)
	DSN         string        // DSN is used to connect to the database
	PoolSize    uint8         // Maximum number of open connections in the pool
	MaxIdle     uint8         // Maximum number of idle connections in the pool
	MaxLifetime time.Duration // Maximum lifetime of a connection before it's recycled
}

func NewDBConfig(driver, dsn string, size, maxIdle uint8, maxLT time.Duration) *DBConfig {
	return &DBConfig{
		Driver:      driver,
		DSN:         dsn,
		PoolSize:    size,
		MaxIdle:     maxIdle,
		MaxLifetime: maxLT,
	}
}

// LoadEnvDBConfig loads the database configuration from environment variables.
// Required variables: DB_DSN, DB_POOL_SIZE, DB_MAX_IDLE, DB_MAX_LIFETIME.
func LoadEnvDBConfig(prefix ...string) (*DBConfig, error) {
	pfx := modPrefix(prefix...)

	driver, err := env.GetEnv(pfx + dbSuffixes[0])
	if err != nil {
		return nil, err
	}

	dsn, err := env.GetEnv(pfx + dbSuffixes[1])
	if err != nil {
		return nil, err
	}

	size, err := env.ParseEnvInt(pfx + dbSuffixes[2])
	if err != nil {
		return nil, err
	}

	maxIdle, err := env.ParseEnvInt(pfx + dbSuffixes[3])
	if err != nil {
		return nil, err
	}

	maxLT, err := env.ParseEnvTimeDuration(pfx + dbSuffixes[4])
	if err != nil {
		return nil, err
	}

	return NewDBConfig(driver, dsn, uint8(size), uint8(maxIdle), maxLT*time.Second), nil
}

// LoadEnvDBConfigs loads multiple DBConfig instances by scanning environment variables
// with the dbSuffixes keys and optional prefixes.
func LoadEnvDBConfigs() (MultiEnvConfig[DBConfig], error) {
	return LoadMultiEnvConfigs(dbSuffixes, LoadEnvDBConfig)
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

// WithDriver sets the Driver string and returns the updated DBConfig.
func (cfg *DBConfig) WithDriver(driver string) *DBConfig {
	cfg.Driver = driver
	return cfg
}
