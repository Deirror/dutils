package appenv

import (
	"github.com/Deirror/servette/env"
)

type MultiEnvEnvConfig = MultiEnvConfig[EnvConfig]

var envSuffixes = []string{"ENV_MODE", "ENV_DOMAIN"}

// EnvConfig holds basic environment configuration like mode and domain.
type EnvConfig struct {
	Mode   string // Application mode: development, staging, production
	Domain string // Public-facing domain, e.g., example.com
}

func NewEnvConfig(mode, domain string) *EnvConfig {
	return &EnvConfig{
		Mode:   mode,
		Domain: domain,
	}
}

// LoadEnvConfig loads EnvConfig from environment variables.
// Required vars: ENV_MODE, ENV_DOMAIN
func LoadEnvConfig(prefix ...string) (*EnvConfig, error) {
	pfx := modPrefix(prefix...)

	mode, err := env.GetEnv(pfx + envSuffixes[0])
	if err != nil {
		return nil, err
	}

	domain, err := env.GetEnv(pfx + envSuffixes[1])
	if err != nil {
		return nil, err
	}

	return NewEnvConfig(mode, domain), nil
}

// LoadEnvConfigs scans env vars and builds env configs based on their prefix.
func LoadEnvConfigs() (MultiEnvConfig[EnvConfig], error) {
	return LoadMultiEnvConfigs(envSuffixes, LoadEnvConfig)
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
