package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/amorindev/headless-ecomm-cms/cmd/api/message"
	"github.com/amorindev/headless-ecomm-cms/cmd/api/middlewares"
	"github.com/amorindev/headless-ecomm-cms/internal/auth"
)

func (h *Handler) SignOut(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	refreshTokenClaims, ok := r.Context().Value(middlewares.RefreshTokenClaimsKey).(*auth.RefreshTokenClaims)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: "Failed to retrieve access token claims"})
		return
	}

	// * Validate
	if refreshTokenClaims.ID == "" {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: "refresh token ID missing in token claims"})
		return
	}

	err := h.AuthMethodSrv.SignOut(context.Background(), refreshTokenClaims.ID)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
