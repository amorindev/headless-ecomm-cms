package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/amorindev/headless-ecomm-cms/cmd/api/message"
	"github.com/amorindev/headless-ecomm-cms/pkg/app/auth-methods/core"
)

func (h Handler) VerifyEmailOtp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req core.VerifyEmailOTPReq

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: "Invalid request format"})
		return
	}

	defer r.Body.Close()

	err = req.IsVerifyEmailOTPValid()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: err.Error()})
		return
	}

	user, session, err := h.AuthMethodSrv.VerifyEmailOtp(context.Background(), req.OtpID, req.OtpCode, req.UserID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: err.Error()})
		return
	}

	resp := core.NewSignInResp(user, session)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
