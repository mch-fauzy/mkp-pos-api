package shared

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/rs/zerolog/log"
)

func SignJWTToken(username, role string, jwtAccessKey []byte) (string, error) {

	// Create a new token with the standard claims and custom claims (username and role)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"role":     role,
		"exp":      time.Now().Add(time.Hour * 2).Unix(),
	})

	// Sign the token with the retrieved secret key
	tokenString, err := token.SignedString(jwtAccessKey)
	if err != nil {
		log.Error().Err(err).Msg("[SignJWTToken] Failed to sign JWT Token")
		return "", err
	}

	return tokenString, nil
}
