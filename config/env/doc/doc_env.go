package docenv

import (
	"github.com/Deirror/servette/env"
	"github.com/Deirror/servette/config/env"
	"github.com/Deirror/servette/domain/doc"
)

type MultiConfig = envcfg.MultiConfig[doc.Config]

var suffixes = []string{
	"DOC_STORE_URL",
}

// LoadConfig loads the document store configuration from environment variables,
// supporting an optional prefix.
func LoadConfig(prefix ...string) (*doc.Config, error) {
	pfx := envcfg.ModPrefix(prefix...)

	url, err := env.Get(pfx + suffixes[0])
	if err != nil {
		return nil, err
	}

	return doc.NewConfig(url), nil
}

// LoadMultiConfigs loads multiple Config instances by scanning env vars with suffixes.
func LoadMultiConfigs() (MultiConfig, error) {
	return envcfg.LoadMultiConfig(suffixes, LoadConfig)
}
