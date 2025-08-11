package handler

import (
	"encoding/json"
	"net/http"

	"github.com/amorindev/headless-ecomm-cms/cmd/api/message"
	"github.com/amorindev/headless-ecomm-cms/cmd/api/middlewares"
	"github.com/amorindev/headless-ecomm-cms/internal/auth"
	"github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/addresses/domain"
)


func (h Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	accessTokenClaim, ok := r.Context().Value(middlewares.AccessTokenClaimsIDKey).(*auth.AccessTokenClaims)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: "Failed to retrieve access token claims"})
		return
	}

	if accessTokenClaim.Subject == "" {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: "User ID missing in token claims"})
		return
	}

	addresses, err := h.AddressSrv.GetAll(r.Context(), accessTokenClaim.UserID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: err.Error()})
		return
	}

	type respBody struct {
		Addresses []*domain.Address `json:"addresses"`
	}

	resp := respBody{
		Addresses: addresses,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
