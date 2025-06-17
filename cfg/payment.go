package cfg

import "github.com/Deirror/dutils/env"

type PaymentConfig struct {
	APIKey        string
	SecretKey     string
	WebhookURL    string
	WebhookSecret string
}

func NewPaymentConfig(apiKey, secretKey, url, webhookSecret string) *PaymentConfig {
	return &PaymentConfig{
		APIKey:        apiKey,
		SecretKey:     secretKey,
		WebhookURL:    url,
		WebhookSecret: webhookSecret,
	}
}

func LoadEnvPaymentConfig() (*PaymentConfig, error) {
	apiKey, err := env.GetEnv("PAYMENT_API_KEY")
	if err != nil {
		return nil, err
	}

	secretKey, err := env.GetEnv("PAYMENT_SECRET_KEY")
	if err != nil {
		return nil, err
	}

	webhookURL, err := env.GetEnv("PAYMENT_WEBHOOK_URL")
	if err != nil {
		return nil, err
	}

	webhookSecret, err := env.GetEnv("PAYMENT_WEBHOOK_SECRET")
	if err != nil {
		return nil, err
	}

	return NewPaymentConfig(apiKey, secretKey, webhookURL, webhookSecret), nil
}

func (cfg *PaymentConfig) WithAPIKey(key string) *PaymentConfig {
	cfg.APIKey = key
	return cfg
}

func (cfg *PaymentConfig) WithSecretKey(secret string) *PaymentConfig {
	cfg.SecretKey = secret
	return cfg
}

func (cfg *PaymentConfig) WithWebhookURL(url string) *PaymentConfig {
	cfg.WebhookURL = url
	return cfg
}

func (cfg *PaymentConfig) WithWebhookSecret(secret string) *PaymentConfig {
	cfg.WebhookSecret = secret
	return cfg
}
