package cfg

import "github.com/Deirror/dutils/env"

// EnvConfig holds basic environment configuration like mode and domain.
type EnvConfig struct {
	Mode   string // Application mode: development, staging, production.
	Domain string // Public-facing domain, e.g., example.com
}

func NewEnvConfig(mode, domain string) *EnvConfig {
	return &EnvConfig{
		Mode:   mode,
		Domain: domain,
	}
}

// LoadEnvConfig loads EnvConfig from environment variables.
// Required vars: MODE, DOMAIN
func LoadEnvConfig() (*EnvConfig, error) {
	mode, err := env.GetEnv("MODE")
	if err != nil {
		return nil, err
	}

	domain, err := env.GetEnv("DOMAIN")
	if err != nil {
		return nil, err
	}

	return NewEnvConfig(mode, domain), nil
}

// WithMode sets the mode and returns the updated EnvConfig.
func (cfg *EnvConfig) WithMode(mode string) *EnvConfig {
	cfg.Mode = mode
	return cfg
}

// WithDomain sets the domain and returns the updated EnvConfig.
func (cfg *EnvConfig) WithDomain(domain string) *EnvConfig {
	cfg.Domain = domain
	return cfg
}
