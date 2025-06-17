package jwt

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Deirror/dutils/cfg"
	"github.com/dgrijalva/jwt-go"
)

type JWT struct {
	CookieName string
	Secret     string
	TokenTTL   time.Duration
}

func NewJWT(jwtCfg *cfg.JWTConfig) *JWT {
	return &JWT{
		CookieName: jwtCfg.CookieName,
		Secret:     jwtCfg.Secret,
		TokenTTL:   jwtCfg.TokenTTL,
	}
}

func (j *JWT) GenerateToken(email string) (string, error) {
	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(j.TokenTTL).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.Secret))
}

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

func (j *JWT) ValidateJWT(tokenString string) (string, error) {
	// Parse the JWT
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return j.Secret, nil
	})

	if err != nil {
		return "", err
	}

	// Validate claims and extract email
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Check expiry
		exp, ok := claims["exp"].(float64)
		if !ok {
			return "", fmt.Errorf("missing exp claim")
		}
		if time.Now().Unix() > int64(exp) {
			return "", fmt.Errorf("token has expired")
		}

		// Extract email claim
		email, ok := claims["email"].(string)
		if !ok || email == "" {
			return "", fmt.Errorf("email claim missing or invalid")
		}

		return email, nil
	}

	return "", fmt.Errorf("invalid token claims")
}
