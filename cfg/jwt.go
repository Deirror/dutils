package cfg

import (
	"time"

	"github.com/Deirror/dutils/env"
)

// JWTConfig holds the configuration related to JWT-based authentication.
type JWTConfig struct {
	CookieName string        // Name of the cookie that stores the JWT
	Secret     string        // Secret key used to sign JWTs
	TokenTTL   time.Duration // Time-to-live duration of the token
}

func NewJWTConfig(name, secret string, ttl time.Duration) *JWTConfig {
	return &JWTConfig{
		CookieName: name,
		Secret:     secret,
		TokenTTL:   ttl,
	}
}

// LoadEnvJWTConfig loads the JWT configuration from environment variables:
// JWT_COOKIE_NAME, JWT_SECRET, and JWT_TOKEN_TTL.
func LoadEnvJWTConfig() (*JWTConfig, error) {
	secret, err := env.GetEnv("JWT_SECRET")
	if err != nil {
		return nil, err
	}

	name, err := env.GetEnv("JWT_COOKIE_NAME")
	if err != nil {
		return nil, err
	}

	ttl, err := env.ParseEnvTimeDuration("JWT_TOKEN_TTL")
	if err != nil {
		return nil, err
	}

	return NewJWTConfig(name, secret, ttl), nil
}

// WithCookieName sets the cookie name for the JWTConfig.
func (cfg *JWTConfig) WithCookieName(name string) *JWTConfig {
	cfg.CookieName = name
	return cfg
}

// WithTokenTTL sets the token time-to-live for the JWTConfig.
func (cfg *JWTConfig) WithTokenTTL(ttl time.Duration) *JWTConfig {
	cfg.TokenTTL = ttl
	return cfg
}

// WithSecret sets the secret key for signing JWTs in the JWTConfig.
func (cfg *JWTConfig) WithSecret(secret string) *JWTConfig {
	cfg.Secret = secret
	return cfg
}
