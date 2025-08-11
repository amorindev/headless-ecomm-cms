package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/amorindev/headless-ecomm-cms/cmd/api/message"
	"github.com/amorindev/headless-ecomm-cms/pkg/app/auth-methods/core"
)

func (h Handler) ResendVerifyEmailOtp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req core.EmailReq

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: "Invalid request format"})
		return
	}

	defer r.Body.Close()

	err = req.IsEmailValid()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: err.Error()})
		return
	}

	otpID, err := h.AuthMethodSrv.ResendVerifyEmailOtp(context.Background(), req.Email)
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
