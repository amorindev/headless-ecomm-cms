package auth

import (
	"errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type RefreshTokenClaims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

func NewRefreshTokenClaim(userID string, rememberMe bool) *RefreshTokenClaims {
	refreshID := uuid.New().String()
	expiresAt := time.Hour * 24 * 7
	if rememberMe {
		expiresAt = time.Hour * 24 * 30
	}
	return &RefreshTokenClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        refreshID,
			Subject:   userID,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiresAt)),
		},
	}
}

func (c *RefreshTokenClaims) GetToken(refreshString string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString([]byte(refreshString))
}

func GetRefreshTokenFromJWT(tokenString, refreshString string) (*RefreshTokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &RefreshTokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(refreshString), nil
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
		return nil, errors.New("token valid: false")
	}

	claim, ok := token.Claims.(*RefreshTokenClaims)
	if !ok {
		return nil, errors.New("invalid claim")
	}

	return claim, nil
}
