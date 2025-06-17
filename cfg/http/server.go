package cfg

import (
	"time"

	"github.com/Deirror/dutils/env"
)

// HTTPServerConfig contains configuration settings for an HTTP server.
type HTTPServerConfig struct {
	Port         string        // Port on which the server listens
	ReadTimeout  time.Duration // Maximum duration for reading the entire request
	WriteTimeout time.Duration // Maximum duration before timing out writes
	IdleTimeout  time.Duration // Maximum time to wait for the next request
}

func NewHTTPServerConfig(port string, read, write, idle time.Duration) *HTTPServerConfig {
	return &HTTPServerConfig{
		Port:         port,
		ReadTimeout:  read,
		WriteTimeout: write,
		IdleTimeout:  idle,
	}
}

// LoadEnvHTTPServerConfig loads HTTP server configuration from environment variables.
// It expects the following environment variables to be set:
//   - HTTPSERVER_PORT
//   - HTTPSERVER_READ_TIMEOUT
//   - HTTPSERVER_WRITE_TIMEOUT
//   - HTTPSERVER_IDLE_TIMEOUT
//
// The timeouts must be in valid Go duration format (e.g., "5s", "1m").
func LoadEnvHTTPServerConfig() (*HTTPServerConfig, error) {
	port, err := env.GetEnv("HTTPSERVER_PORT")
	if err != nil {
		return nil, err
	}

	readTimeout, err := env.ParseEnvTimeDuration("HTTPSERVER_READ_TIMEOUT")
	if err != nil {
		return nil, err
	}

	writeTimeout, err := env.ParseEnvTimeDuration("HTTPSERVER_WRITE_TIMEOUT")
	if err != nil {
		return nil, err
	}

	idleTimeout, err := env.ParseEnvTimeDuration("HTTPSERVER_IDLE_TIMEOUT")
	if err != nil {
		return nil, err
	}

	return NewHTTPServerConfig(port, readTimeout, writeTimeout, idleTimeout), nil
}

// WithPort sets the Port field and returns the config.
func (cfg *HTTPServerConfig) WithPort(port string) *HTTPServerConfig {
	cfg.Port = port
	return cfg
}

// WithReadTimeout sets the ReadTimeout field and returns the config.
func (cfg *HTTPServerConfig) WithReadTimeout(timeout time.Duration) *HTTPServerConfig {
	cfg.ReadTimeout = timeout
	return cfg
}

// WithWriteTimeout sets the WriteTimeout field and returns the config.
func (cfg *HTTPServerConfig) WithWriteTimeout(timeout time.Duration) *HTTPServerConfig {
	cfg.WriteTimeout = timeout
	return cfg
}

// WithIdleTimeout sets the IdleTimeout field and returns the config.
func (cfg *HTTPServerConfig) WithIdleTimeout(timeout time.Duration) *HTTPServerConfig {
	cfg.IdleTimeout = timeout
	return cfg
}
