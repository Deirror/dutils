package oauth

import (
	"github.com/Deirror/dutils/cfg"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// Wrapper for a ready google oauth config
func NewOAuth2Config(cfg *cfg.OAuthConfig) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
		RedirectURL:  cfg.RedirectURL,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}
}
