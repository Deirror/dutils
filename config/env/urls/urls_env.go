package cfg

import "github.com/Deirror/dutils/env"

type MultiEnvExternalURLConfig = MultiEnvConfig[ExternalURLConfig]

// ExternalURLConfig holds configuration for an external URL.
type ExternalURLConfig struct {
	URL string // External service URL
}

func NewExternalURLConfig(url string) *ExternalURLConfig {
	return &ExternalURLConfig{
		URL: url,
	}
}

// externalURLSuffixes defines environment variable suffixes for ExternalURLConfig.
var externalURLSuffixes = []string{
	"EXTERNAL_URL",
}

// LoadEnvExternalURLConfig loads ExternalURLConfig from environment variables.
// Optionally accepts a prefix to prepend to suffixes.
func LoadEnvExternalURLConfig(prefix ...string) (*ExternalURLConfig, error) {
	pfx := modPrefix(prefix...)

	url, err := env.GetEnv(pfx + externalURLSuffixes[0])
	if err != nil {
		return nil, err
	}

	return NewExternalURLConfig(url), nil
}

// LoadEnvExternalURLConfigs loads multiple ExternalURLConfig from environment variables
// by scanning for all environment variable sets matching externalURLSuffixes.
func LoadEnvExternalURLConfigs() (MultiEnvConfig[ExternalURLConfig], error) {
	return LoadMultiEnvConfigs(externalURLSuffixes, LoadEnvExternalURLConfig)
}

// WithURL sets the URL field and returns the updated config.
func (cfg *ExternalURLConfig) WithURL(url string) *ExternalURLConfig {
	cfg.URL = url
	return cfg
}
