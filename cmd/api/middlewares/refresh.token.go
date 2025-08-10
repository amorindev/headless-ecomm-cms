package middlewares

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/amorindev/headless-ecomm-cms/cmd/api/message"
)

func (m *AuthMiddleware) RefreshTokenMdw(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			RefreshToken string `json:"refresh_token"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		if req.RefreshToken == "" {
			http.Error(w, "Refresh token required", http.StatusBadRequest)
			return
		}

		c, err := m.AuthSrv.ParseRefreshToken(req.RefreshToken)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(message.ErrorMessage{Message: err.Error()})
			return
		}

		ctx := context.WithValue(r.Context(), RefreshTokenClaimsKey, c)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
