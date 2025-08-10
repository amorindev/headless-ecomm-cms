package middlewares

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/amorindev/headless-ecomm-cms/cmd/api/message"
)

func (m *AuthMiddleware) AccessTokenMdw(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		tokenString, err := tokenFromAuthorization(authHeader)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(message.ErrorMessage{Message: err.Error()})
			return
		}
		c, err := m.AuthSrv.ParseAccessToken(tokenString)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(message.ErrorMessage{Message: err.Error()})
			return
		}
		ctx := context.WithValue(r.Context(), UserIDKey, c.Subject)
		ctx = context.WithValue(ctx, AccessTokenClaimsIDKey, c)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

func tokenFromAuthorization(authorization string) (string, error) {
	if authorization == "" {
		return "", errors.New("authorization header is required")
	}

	if !strings.HasPrefix(authorization, "Bearer") {
		return "", errors.New("invalid authorization format")
	}

	l := strings.Split(authorization, " ")
	if len(l) != 2 {
		return "", errors.New("invalid authorization format")
	}

	return l[1], nil
}
