package cfg

import (
	"time"

	"github.com/Deirror/dutils/env"
)

var serverSuffixes = []string{
	"SERVER_PORT",
	"SERVER_READ_TIMEOUT",
	"SERVER_WRITE_TIMEOUT",
	"SERVER_IDLE_TIMEOUT",
}

// ServerConfig contains configuration settings for an server.
type ServerConfig struct {
	Port         string        // Port on which the server listens
	ReadTimeout  time.Duration // Maximum duration for reading the entire request
	WriteTimeout time.Duration // Maximum duration before timing out writes
	IdleTimeout  time.Duration // Maximum time to wait for the next request
}

func NewServerConfig(port string, read, write, idle time.Duration) *ServerConfig {
	return &ServerConfig{
		Port:         port,
		ReadTimeout:  read,
		WriteTimeout: write,
		IdleTimeout:  idle,
	}
}

// LoadEnvServerConfig loads server config values from environment variables.
// The env var keys are prefixed with the optional prefix argument.
func LoadEnvServerConfig(prefix ...string) (*ServerConfig, error) {
	pfx := modPrefix(prefix...)

	port, err := env.GetEnv(pfx + serverSuffixes[0])
	if err != nil {
		return nil, err
	}

	readTimeout, err := env.ParseEnvTimeDuration(pfx + serverSuffixes[1])
	if err != nil {
		return nil, err
	}

	writeTimeout, err := env.ParseEnvTimeDuration(pfx + serverSuffixes[2])
	if err != nil {
		return nil, err
	}

	idleTimeout, err := env.ParseEnvTimeDuration(pfx + serverSuffixes[3])
	if err != nil {
		return nil, err
	}

	return NewServerConfig(port, readTimeout, writeTimeout, idleTimeout), nil
}

// LoadEnvServerConfigs scans env vars and builds Server configs based on their prefix.
func LoadEnvServerConfigs() (MultiEnvConfig[ServerConfig], error) {
	return LoadMultiEnvConfigs(serverSuffixes, LoadEnvServerConfig)
}

// WithPort sets the Port field and returns the updated ServerConfig.
func (cfg *ServerConfig) WithPort(port string) *ServerConfig {
	cfg.Port = port
	return cfg
}

// WithReadTimeout sets the ReadTimeout field and returns the updated ServerConfig.
func (cfg *ServerConfig) WithReadTimeout(timeout time.Duration) *ServerConfig {
	cfg.ReadTimeout = timeout
	return cfg
}

// WithWriteTimeout sets the WriteTimeout field and returns the updated ServerConfig.
func (cfg *ServerConfig) WithWriteTimeout(timeout time.Duration) *ServerConfig {
	cfg.WriteTimeout = timeout
	return cfg
}

// WithIdleTimeout sets the IdleTimeout field and returns the updated ServerConfig.
func (cfg *ServerConfig) WithIdleTimeout(timeout time.Duration) *ServerConfig {
	cfg.IdleTimeout = timeout
	return cfg
}
