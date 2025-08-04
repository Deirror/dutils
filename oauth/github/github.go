package github

import (
	"github.com/Deirror/dutils/cfg"
	"github.com/Deirror/dutils/oauth"
	"golang.org/x/oauth2"
)

var Scopes = []string{"read:user", "user:email"}

var Endpoint = oauth2.Endpoint{
	AuthURL:  "https://github.com/login/oauth/authorize",
	TokenURL: "https://github.com/login/oauth/access_token",
}

func NewOAuth2Config(cfg *cfg.OAuthConfig) *oauth2.Config {
	return oauth.NewOAuth2Config(
		cfg,
		Scopes,
		Endpoint,
	)
}
