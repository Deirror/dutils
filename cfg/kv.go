package cfg

import "github.com/Deirror/dutils/env"

var kvSuffixes = []string{
	"KV_STORE_URL",
}

// KVConfig holds configuration details for connecting to a key-value store.
type KVConfig struct {
	KVStoreURL string // URL of the key-value store
}

func NewKVConfig(kvStoreURL string) *KVConfig {
	return &KVConfig{
		KVStoreURL: kvStoreURL,
	}
}

// LoadEnvKVConfig loads the key-value store configuration from environment variables,
// supporting an optional prefix.
func LoadEnvKVConfig(prefix ...string) (*KVConfig, error) {
	pfx := modPrefix(prefix...)

	url, err := env.GetEnv(pfx + kvSuffixes[0])
	if err != nil {
		return nil, err
	}

	return NewKVConfig(url), nil
}

// LoadEnvKVConfigs loads multiple KVConfig instances by scanning env vars with kvSuffixes.
func LoadEnvKVConfigs() (MultiEnvConfig[KVConfig], error) {
	return LoadMultiEnvConfigs(kvSuffixes, LoadEnvKVConfig)
}

// WithKVStoreURL sets a new key-value store URL and returns the updated KVConfig.
func (cfg *KVConfig) WithKVStoreURL(url string) *KVConfig {
	cfg.KVStoreURL = url
	return cfg
}
