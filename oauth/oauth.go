package oauth

import (
	"github.com/Deirror/dutils/cfg"
	"golang.org/x/oauth2"
)

func NewOAuth2Config(cfg *cfg.OAuthConfig, scopes []string, urls oauth2.Endpoint) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
		RedirectURL:  cfg.RedirectURL,
		Scopes:       scopes,
		Endpoint:     urls,
	}
}
