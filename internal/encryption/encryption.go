package encryption

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword takes a plain-text password and returns its hashed version
// using bcrypt. It uses bcrypt.DefaultCost as the cost factor.
// Returns the hash as a string and an error if hashing fails.
func HashPassword(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password %w", err)
	}
	return string(hashPassword), nil
}

// CheckPassword compares a plain-text password with its hashed value.
// Returns nil if they match, or an error if they do not match or if
// verification fails.
func CheckPassword(password string, hashPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
}
