package jwt

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Deirror/dutils/cfg"
	"github.com/dgrijalva/jwt-go"
)

// JWT handles JWT token creation, validation, and cookie management.
type JWT struct {
	CookieName string        // Name of the cookie to store the JWT
	Secret     string        // Secret key used to sign the JWT
	TokenTTL   time.Duration // Token time-to-live duration
}

func NewJWT(jwtCfg *cfg.JWTConfig) *JWT {
	return &JWT{
		CookieName: jwtCfg.CookieName,
		Secret:     jwtCfg.Secret,
		TokenTTL:   jwtCfg.TokenTTL,
	}
}

// GenerateToken creates a signed JWT token containing the user's email and expiration.
func (j *JWT) GenerateToken(email string) (string, error) {
	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(j.TokenTTL).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.Secret))
}

// SetCookie sets an HTTP-only, secure cookie with the JWT token on the response writer.
func (j *JWT) SetCookie(w http.ResponseWriter, token string) {
	cookie := &http.Cookie{
		Name:     j.CookieName,
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		Expires:  time.Now().Add(j.TokenTTL),
	}
	http.SetCookie(w, cookie)
}

// ValidateJWT parses and validates the JWT token string.
// Returns the email claim if valid, or an error otherwise.
func (j *JWT) ValidateJWT(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(j.Secret), nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		exp, ok := claims["exp"].(float64)
		if !ok {
			return "", fmt.Errorf("missing exp claim")
		}
		if time.Now().Unix() > int64(exp) {
			return "", fmt.Errorf("token has expired")
		}

		email, ok := claims["email"].(string)
		if !ok || email == "" {
			return "", fmt.Errorf("email claim missing or invalid")
		}

		return email, nil
	}

	return "", fmt.Errorf("invalid token claims")
}
