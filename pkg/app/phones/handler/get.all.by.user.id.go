package handler

import (
	"encoding/json"
	"net/http"

	"github.com/amorindev/headless-ecomm-cms/cmd/api/message"
	md "github.com/amorindev/headless-ecomm-cms/cmd/api/middlewares"
	"github.com/amorindev/headless-ecomm-cms/internal/auth"
	"github.com/amorindev/headless-ecomm-cms/pkg/app/phones/domain"
)

func (h Handler) GetAllByUserID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	accessTokenClaim, ok := r.Context().Value(md.AccessTokenClaimsIDKey).(*auth.AccessTokenClaims)
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

	phones, err := h.PhoneSrv.GetAllByUserID(r.Context(), accessTokenClaim.UserID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: err.Error()})
		return
	}

	type respBody struct {
		Phones []*domain.Phone `json:"phones"`
	}

	resp := respBody{
		Phones: phones,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
