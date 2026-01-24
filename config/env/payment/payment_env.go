package cfg

import "github.com/Deirror/dutils/env"

type MultiEnvPaymentConfig = MultiEnvConfig[PaymentConfig]

var paymentSuffixes = []string{
	"PAYMENT_API_KEY",
	"PAYMENT_SECRET_KEY",
	"PAYMENT_WEBHOOK_URL",
	"PAYMENT_WEBHOOK_SECRET",
}

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

// LoadEnvPaymentConfig loads PaymentConfig from environment variables,
// optionally with a prefix.
func LoadEnvPaymentConfig(prefix ...string) (*PaymentConfig, error) {
	pfx := modPrefix(prefix...)

	apiKey, err := env.GetEnv(pfx + paymentSuffixes[0])
	if err != nil {
		return nil, err
	}

	secretKey, err := env.GetEnv(pfx + paymentSuffixes[1])
	if err != nil {
		return nil, err
	}

	webhookURL, err := env.GetEnv(pfx + paymentSuffixes[2])
	if err != nil {
		return nil, err
	}

	webhookSecret, err := env.GetEnv(pfx + paymentSuffixes[3])
	if err != nil {
		return nil, err
	}

	return NewPaymentConfig(apiKey, secretKey, webhookURL, webhookSecret), nil
}

// LoadEnvPaymentConfigs loads multiple PaymentConfigs by scanning env vars with payment suffixes.
func LoadEnvPaymentConfigs() (MultiEnvConfig[PaymentConfig], error) {
	return LoadMultiEnvConfigs(paymentSuffixes, LoadEnvPaymentConfig)
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
