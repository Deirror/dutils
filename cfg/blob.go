package cfg

import "github.com/Deirror/dutils/env"

type BlobConfig struct {
	ProjectURL string
	APIKey     string
	Bucket     string
}

func NewBlobConfig(url, apiKey, bucket string) *BlobConfig {
	return &BlobConfig{
		ProjectURL: url,
		APIKey:     apiKey,
		Bucket:     bucket,
	}
}

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

func (cfg *BlobConfig) WithProjectURL(url string) *BlobConfig {
	cfg.ProjectURL = url
	return cfg
}

func (cfg *BlobConfig) WithAPIKey(apiKey string) *BlobConfig {
	cfg.ProjectURL = apiKey
	return cfg
}

func (cfg *BlobConfig) WithBucket(bucket string) *BlobConfig {
	cfg.ProjectURL = bucket
	return cfg
}
