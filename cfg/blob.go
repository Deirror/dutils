package cfg

import "github.com/Deirror/dutils/env"

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

// LoadEnvBlobConfig reads the BlobConfig values from environment variables:
// BLOB_PROJECT_URL, BLOB_API_KEY, and BLOB_BUCKET.
func LoadEnvBlobConfig() (*BlobConfig, error) {
	url, err := env.GetEnv("BLOB_PROJECT_URL")
	if err != nil {
		return nil, err
	}

	apiKey, err := env.GetEnv("BLOB_API_KEY")
	if err != nil {
		return nil, err
	}

	bucket, err := env.GetEnv("BLOB_BUCKET")
	if err != nil {
		return nil, err
	}

	return NewBlobConfig(url, apiKey, bucket), nil
}

// WithProjectURL sets the ProjectURL and returns the updated BlobConfig.
func (cfg *BlobConfig) WithProjectURL(url string) *BlobConfig {
	cfg.ProjectURL = url
	return cfg
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
