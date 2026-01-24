package appenv

import (
	"github.com/Deirror/servette/app"
	"github.com/Deirror/servette/config/env"
	"github.com/Deirror/servette/env"
)

type MultiConfig = envcfg.MultiConfig[app.Config]

var suffixes = []string{"APP_MODE", "APP_DOMAIN"}

// LoadConfig loads Config from environment variables.
// Required vars: APP_MODE, APP_DOMAIN
func LoadConfig(prefix ...string) (*app.Config, error) {
	pfx := envcfg.ModPrefix(prefix...)

	mode, err := env.Get(pfx + suffixes[0])
	if err != nil {
		return nil, err
	}

	domain, err := env.Get(pfx + suffixes[1])
	if err != nil {
		return nil, err
	}

	return app.NewConfig(mode, domain), nil
}

// LoadMultiConfig scans env vars and builds app configs based on their prefix.
func LoadMultiConfig() (MultiConfig, error) {
	return envcfg.LoadMultiConfig(suffixes, LoadConfig)
}
