package oauth

// UserInfoURLs maps OAuth provider names to their user info API endpoints.
// These URLs are used to retrieve basic profile data (such as ID, name, and email)
// after exchanging an access token during the OAuth2 flow.
var UserInfoURLs = map[string]string{
	"google":   "https://www.googleapis.com/oauth2/v2/userinfo",
	"facebook": "https://graph.facebook.com/me?fields=id,name,email,picture",
	"github":   "https://api.github.com/user",
}
