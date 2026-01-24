package cfg

import "github.com/Deirror/dutils/env"

type MultiEnvOAuthConfig = MultiEnvConfig[OAuthConfig]

var oauthSuffixes = []string{
	"OAUTH_CLIENT_ID",
	"OAUTH_CLIENT_SECRET",
	"OAUTH_REDIRECT_URL",
}

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

// LoadEnvOAuthConfig loads OAuthConfig from environment variables,
// with an optional prefix.
func LoadEnvOAuthConfig(prefix ...string) (*OAuthConfig, error) {
	pfx := modPrefix(prefix...)

	clientID, err := env.GetEnv(pfx + oauthSuffixes[0])
	if err != nil {
		return nil, err
	}

	clientSecret, err := env.GetEnv(pfx + oauthSuffixes[1])
	if err != nil {
		return nil, err
	}

	redirectURL, err := env.GetEnv(pfx + oauthSuffixes[2])
	if err != nil {
		return nil, err
	}

	return NewOAuthConfig(clientID, clientSecret, redirectURL), nil
}

// LoadEnvOAuthConfigs loads multiple OAuthConfigs by scanning env vars with oauth suffixes.
func LoadEnvOAuthConfigs() (MultiEnvConfig[OAuthConfig], error) {
	return LoadMultiEnvConfigs(oauthSuffixes, LoadEnvOAuthConfig)
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
