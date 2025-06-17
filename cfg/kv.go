package cfg

import "github.com/Deirror/dutils/env"

type KVConfig struct {
	KVStoreURL string
}

func NewKVConfig(kvStoreURL string) *KVConfig {
	return &KVConfig{
		KVStoreURL: kvStoreURL,
	}
}

func LoadEnvKVConfig() (*KVConfig, error) {
	url, err := env.GetEnv("KV_STORE_URL")
	if err != nil {
		return nil, err
	}

	return NewKVConfig(url), nil
}

func (cfg *KVConfig) WithKVStoreURL(url string) *KVConfig {
	cfg.KVStoreURL = url
	return cfg
}
