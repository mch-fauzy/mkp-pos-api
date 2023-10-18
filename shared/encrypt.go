package shared

import (
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes a password using bcrypt.
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Error().Err(err).Msg("[HashPassword] Error hashing password")
		return "", err
	}
	return string(hashedPassword), nil
}

// VerifyPassword checks if a password matches a hashed password.
func VerifyPassword(password, hashedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		log.Error().Err(err).Msg("[VerifyPassword] Error verifying password")
		return err
	}
	return nil
}
