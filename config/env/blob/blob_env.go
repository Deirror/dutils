package blobenv

import (
	"github.com/Deirror/servette/config/env"
	"github.com/Deirror/servette/domain/blob"
	"github.com/Deirror/servette/env"
)

type MultiEnvBlobConfig = envcfg.MultiEnvConfig[blob.Config]

var blobSuffixes = []string{"BLOB_PROJECT_URL", "BLOB_API_KEY", "BLOB_BUCKET"}

// LoadEnvBlobConfig loads BlobConfig from env vars with optional prefix.
func LoadConfig(prefix ...string) (*blob.Config, error) {
	pfx := envcfg.ModPrefix(prefix...)

	url, err := env.GetEnv(pfx + blobSuffixes[0])
	if err != nil {
		return nil, err
	}

	apiKey, err := env.GetEnv(pfx + blobSuffixes[1])
	if err != nil {
		return nil, err
	}

	bucket, err := env.GetEnv(pfx + blobSuffixes[2])
	if err != nil {
		return nil, err
	}

	return blob.NewConfig(url, apiKey, bucket), nil
}

// LoadEnvBlobConfigs loads multiple BlobConfigs by scanning env vars with blob suffixes.
func LoadEnvBlobConfigs() (MultiEnvConfig[BlobConfig], error) {
	return LoadMultiEnvConfigs(blobSuffixes, LoadEnvBlobConfig)
}

// WithAPIKey sets the APIKey and returns the updated BlobConfig.
func (cfg *BlobConfig) WithAPIKey(apiKey string) *BlobConfig {
	cfg.APIKey = apiKey
	return cfg
}

// WithBucket sets the Bucket and returns the updated BlobConfig.
func (cfg *BlobConfig) WithBucket(bucket string) *BlobConfig {
	cfg.Bucket = bucket
	return cfg
}
