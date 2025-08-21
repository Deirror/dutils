package cfg

import "github.com/Deirror/dutils/env"

type MultiEnvDocConfig = MultiEnvConfig[DocConfig]

var docSuffixes = []string{
	"DOC_STORE_URL",
}

// DocConfig holds configuration details for connecting to a document database.
type DocConfig struct {
	StoreURL string // used such as a MongoDB connection string (example: mongodb+srv://.../dbName)
}

func NewDocConfig(docStoreURL string) *DocConfig {
	return &DocConfig{
		StoreURL: docStoreURL,
	}
}

// LoadEnvDocConfig loads the document store configuration from environment variables,
// supporting an optional prefix.
func LoadEnvDocConfig(prefix ...string) (*DocConfig, error) {
	pfx := modPrefix(prefix...)

	url, err := env.GetEnv(pfx + docSuffixes[0])
	if err != nil {
		return nil, err
	}

	return NewDocConfig(url), nil
}

// LoadEnvDocConfigs loads multiple DocConfig instances by scanning env vars with docSuffixes.
func LoadEnvDocConfigs() (MultiEnvConfig[DocConfig], error) {
	return LoadMultiEnvConfigs(docSuffixes, LoadEnvDocConfig)
}

// WithDocStoreURL sets a new document store URL and returns the updated DocConfig.
func (cfg *DocConfig) WithDocStoreURL(url string) *DocConfig {
	cfg.StoreURL = url
	return cfg
}
