package cfg

import (
	"time"

	"github.com/Deirror/dutils/env"
)

type JWTConfig struct {
	CookieName string
	Secret     string
	TokenTTL   time.Duration
}

func NewJWTConfig(name, secret string, ttl time.Duration) *JWTConfig {
	return &JWTConfig{
		CookieName: name,
		Secret:     secret,
		TokenTTL:   ttl,
	}
}

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

func (cfg *JWTConfig) WithCookieName(name string) *JWTConfig {
	cfg.CookieName = name
	return cfg
}

func (cfg *JWTConfig) WithTokenTTL(ttl time.Duration) *JWTConfig {
	cfg.TokenTTL = ttl
	return cfg
}

func (cfg *JWTConfig) WithSecret(secret string) *JWTConfig {
	cfg.Secret = secret
	return cfg
}
