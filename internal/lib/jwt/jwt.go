package jwt

import (
	"time"

	"github.com/DaniilZ77/authorization/internal/core"
	jwtlib "github.com/golang-jwt/jwt"
)

func GenerateToken(id int, authConfig core.AuthConfig) (*string, error) {
	token := jwtlib.NewWithClaims(jwtlib.SigningMethodHS256, jwtlib.MapClaims{
		"id":  id,
		"iat": time.Now().Add(time.Minute * time.Duration(authConfig.TokenTTL)).Unix(),
	})

	tokenString, err := token.SignedString([]byte(authConfig.Secret))
	if err != nil {
		return nil, err
	}

	return &tokenString, nil
}
