package google

import (
	"github.com/Deirror/servette/config"
	"github.com/Deirror/servette/auth/oauth"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var Scopes = []string{
	"https://www.googleapis.com/auth/userinfo.email",
	"https://www.googleapis.com/auth/userinfo.profile",
}

func NewOAuth2Config(cfg *cfg.OAuthConfig) *oauth2.Config {
	return oauth.NewOAuth2Config(
		cfg,
		Scopes,
		google.Endpoint,
	)
}
