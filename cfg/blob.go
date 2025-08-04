package cfg

import "github.com/Deirror/dutils/env"

var blobSuffixes = []string{"BLOB_PROJECT_URL", "BLOB_API_KEY", "BLOB_BUCKET"}

// BlobConfig holds the configuration required to connect to a blob storage provider.
type BlobConfig struct {
	ProjectURL string // Base project URL of the blob storage service
	APIKey     string // API key for authentication with the blob storage provider
	Bucket     string // Target bucket name for storing or retrieving objects
}

func NewBlobConfig(url, apiKey, bucket string) *BlobConfig {
	return &BlobConfig{
		ProjectURL: url,
		APIKey:     apiKey,
		Bucket:     bucket,
	}
}

// LoadEnvBlobConfig loads BlobConfig from env vars with optional prefix.
func LoadEnvBlobConfig(prefix ...string) (*BlobConfig, error) {
	pfx := modPrefix(prefix...)

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

	return NewBlobConfig(url, apiKey, bucket), nil
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
