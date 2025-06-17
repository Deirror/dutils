package cfg

import "github.com/Deirror/dutils/env"

// PaymentConfig holds configuration details for payment processing.
type PaymentConfig struct {
	APIKey        string // Public API key for the payment provider
	SecretKey     string // Secret key for authenticating requests
	WebhookURL    string // URL to receive payment provider webhook callbacks
	WebhookSecret string // Secret used to verify webhook authenticity
}

func NewPaymentConfig(apiKey, secretKey, url, webhookSecret string) *PaymentConfig {
	return &PaymentConfig{
		APIKey:        apiKey,
		SecretKey:     secretKey,
		WebhookURL:    url,
		WebhookSecret: webhookSecret,
	}
}

// LoadEnvPaymentConfig loads PaymentConfig from environment variables:
// PAYMENT_API_KEY, PAYMENT_SECRET_KEY, PAYMENT_WEBHOOK_URL, PAYMENT_WEBHOOK_SECRET.
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

// WithAPIKey sets the API key and returns the updated config.
func (cfg *PaymentConfig) WithAPIKey(key string) *PaymentConfig {
	cfg.APIKey = key
	return cfg
}

// WithSecretKey sets the secret key and returns the updated config.
func (cfg *PaymentConfig) WithSecretKey(secret string) *PaymentConfig {
	cfg.SecretKey = secret
	return cfg
}

// WithWebhookURL sets the webhook URL and returns the updated config.
func (cfg *PaymentConfig) WithWebhookURL(url string) *PaymentConfig {
	cfg.WebhookURL = url
	return cfg
}

// WithWebhookSecret sets the webhook secret and returns the updated config.
func (cfg *PaymentConfig) WithWebhookSecret(secret string) *PaymentConfig {
	cfg.WebhookSecret = secret
	return cfg
}
