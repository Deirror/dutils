package cfg

import "github.com/Deirror/dutils/env"

// KVConfig holds configuration details for connecting to a key-value store.
type KVConfig struct {
	KVStoreURL string // URL of the key-value store
}

func NewKVConfig(kvStoreURL string) *KVConfig {
	return &KVConfig{
		KVStoreURL: kvStoreURL,
	}
}

// LoadEnvKVConfig loads the key-value store configuration from the environment variable KV_STORE_URL.
func LoadEnvKVConfig() (*KVConfig, error) {
	url, err := env.GetEnv("KV_STORE_URL")
	if err != nil {
		return nil, err
	}

	return NewKVConfig(url), nil
}

// WithKVStoreURL sets a new key-value store URL and returns the updated KVConfig.
func (cfg *KVConfig) WithKVStoreURL(url string) *KVConfig {
	cfg.KVStoreURL = url
	return cfg
}
