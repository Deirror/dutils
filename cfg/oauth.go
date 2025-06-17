package cfg

import "github.com/Deirror/dutils/env"

type OAuthConfig struct {
	ClientID     string
	ClientSecret string
	RedirectURL  string
}

func NewOAuthConfig(clientID, clientSecret, redirectURL string) *OAuthConfig {
	return &OAuthConfig{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
	}
}

func LoadEnvOAuthConfig() (*OAuthConfig, error) {
	clientID, err := env.GetEnv("OAUTH_CLIENT_ID")
	if err != nil {
		return nil, err
	}

	clientSecret, err := env.GetEnv("OAUTH_CLIENT_SECRET")
	if err != nil {
		return nil, err
	}

	redirectURL, err := env.GetEnv("OAUTH_REDIRECT_URL")
	if err != nil {
		return nil, err
	}

	return NewOAuthConfig(clientID, clientSecret, redirectURL), nil
}

func (cfg *OAuthConfig) WithClientID(clientID string) *OAuthConfig {
	cfg.ClientID = clientID
	return cfg
}

func (cfg *OAuthConfig) WithClientSecret(clientSecret string) *OAuthConfig {
	cfg.ClientSecret = clientSecret
	return cfg
}

func (cfg *OAuthConfig) WithRedirectURL(redirectURL string) *OAuthConfig {
	cfg.RedirectURL = redirectURL
	return cfg
}
