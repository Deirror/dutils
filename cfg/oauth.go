package cfg

import "github.com/Deirror/dutils/env"

// OAuthConfig holds OAuth 2.0 client credentials and redirect URL.
type OAuthConfig struct {
	ClientID     string // OAuth client ID
	ClientSecret string // OAuth client secret
	RedirectURL  string // OAuth redirect URL after authentication
}

func NewOAuthConfig(clientID, clientSecret, redirectURL string) *OAuthConfig {
	return &OAuthConfig{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
	}
}

// LoadEnvOAuthConfig loads OAuthConfig from environment variables:
// OAUTH_CLIENT_ID, OAUTH_CLIENT_SECRET, and OAUTH_REDIRECT_URL.
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

// WithClientID sets the OAuth client ID and returns the updated config.
func (cfg *OAuthConfig) WithClientID(clientID string) *OAuthConfig {
	cfg.ClientID = clientID
	return cfg
}

// WithClientSecret sets the OAuth client secret and returns the updated config.
func (cfg *OAuthConfig) WithClientSecret(clientSecret string) *OAuthConfig {
	cfg.ClientSecret = clientSecret
	return cfg
}

// WithRedirectURL sets the OAuth redirect URL and returns the updated config.
func (cfg *OAuthConfig) WithRedirectURL(redirectURL string) *OAuthConfig {
	cfg.RedirectURL = redirectURL
	return cfg
}
