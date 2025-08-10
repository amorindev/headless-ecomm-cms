package auth

import (
	"errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AccessTokenClaims struct {
	UserID string   `json:"user_id"`
	Email  string   `json:"email"`
	Roles  []string `json:"role,omitempty"`
	jwt.RegisteredClaims
}

// NewAccessToken creates a JWT access token
func NewAccessTokenClaim(userID string, email string, issuer string, roles []string) *AccessTokenClaims {
	return &AccessTokenClaims{
		UserID: userID,
		Email:  email,
		Roles:  roles,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   userID,
			Issuer:    issuer,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 2)),
		},
	}
}

func (c *AccessTokenClaims) GetToken(accessString string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString([]byte(accessString))
}

// ValidateToken validates and parses a JWT token
func GetAccessTokenFromJWT(tokenString, signingString string) (*AccessTokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &AccessTokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(signingString), nil
	})

	if err != nil {
		if err.Error() == "token has invalid claims: token is expired" {
			return nil, errors.New("token-is-expired")
		}

		if strings.Contains(err.Error(), "signature is invalid") {
			return nil, errors.New("invalid token signature")
		}

		return nil, errors.New("invalid token format")
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	claim, ok := token.Claims.(*AccessTokenClaims)
	if !ok {
		return nil, errors.New("invalid claims")
	}

	return claim, nil
}
