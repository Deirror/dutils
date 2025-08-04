package facebook

import (
	"github.com/Deirror/dutils/cfg"
	"github.com/Deirror/dutils/oauth"
	"golang.org/x/oauth2"
)

var Scopes = []string{
	"email",
	"public_profile",
}

var Endpoint = oauth2.Endpoint{
	AuthURL:  "https://www.facebook.com/v17.0/dialog/oauth",
	TokenURL: "https://graph.facebook.com/v17.0/oauth/access_token",
}

func NewOAuth2Config(cfg *cfg.OAuthConfig) *oauth2.Config {
	return oauth.NewOAuth2Config(
		cfg,
		Scopes,
		Endpoint,
	)
}
