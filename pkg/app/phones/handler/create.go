package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/amorindev/headless-ecomm-cms/cmd/api/message"
	"github.com/amorindev/headless-ecomm-cms/cmd/api/middlewares"
	"github.com/amorindev/headless-ecomm-cms/internal/auth"
	"github.com/amorindev/headless-ecomm-cms/pkg/app/phones/core"
	"github.com/amorindev/headless-ecomm-cms/pkg/app/phones/domain"
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

	var req core.CreatePhoneReq

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: "Invalid request format"})
		return
	}

	defer r.Body.Close()

	if req.Number == "" || req.CountryCode == "" || req.CountryIso == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: "number, countryCode countryIsoCode are required"})
		return
	}

	phone := &domain.Phone{
		UserID:      accessTokenClaim.Subject,
		Number:      req.Number,
		CountryCode: req.CountryCode,
		CountryIso:  req.CountryIso,
	}

	err = h.PhoneSrv.Create(context.Background(), phone)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(phone)
}
