package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/amorindev/headless-ecomm-cms/cmd/api/message"
	"github.com/amorindev/headless-ecomm-cms/cmd/api/middlewares"
	"github.com/amorindev/headless-ecomm-cms/internal/auth"
	"github.com/amorindev/headless-ecomm-cms/pkg/app/auth-methods/core"
)

func (h Handler) ResendVerifyEnableMfaSmsOtp(w http.ResponseWriter, r *http.Request) {
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

	otpID, err := h.AuthMethodSrv.ResendVerifyEnableMfaSmsOtp(context.TODO(), accessTokenClaim.Subject)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: err.Error()})
		return
	}

	resp := core.OtpIDResp{
		OtpID: otpID,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
