package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/amorindev/headless-ecomm-cms/cmd/api/message"
	"github.com/amorindev/headless-ecomm-cms/cmd/api/middlewares"
	"github.com/amorindev/headless-ecomm-cms/internal/auth"
	"github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/addresses/core"
	"github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/addresses/domain"
)

func (h Handler) Create(w http.ResponseWriter, r *http.Request) {
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

	var req core.CreateAddressReq

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: "Invalid request format"})
		return
	}

	defer r.Body.Close()

	err = req.IsCreateAddressValid()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: err.Error()})
		return
	}

	address := &domain.Address{
		UserID:      accessTokenClaim.Subject,
		Label:       req.Label,
		AddressLine: req.AddressLine,
		State:       req.State,
		Country:     req.Country,
		City:        req.City,
		Latitude:    req.Latitude,
		Longitude:   req.Longitude,
		PostalCode:  req.PostalCode,
	}

	err = h.AddressSrv.Create(context.Background(), address)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(address)
}
