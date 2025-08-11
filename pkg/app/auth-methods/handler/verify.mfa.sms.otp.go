package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/amorindev/headless-ecomm-cms/cmd/api/message"
	"github.com/amorindev/headless-ecomm-cms/pkg/app/auth-methods/core"
)

// Este handler es cuando el usuario inicia session y activo el
// en este punto no tnedre el access token
// asi que aqui no tendre access token ver
// es cuando inicia session
func (h Handler) VerifyMfaSmsOtp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req core.OtpIDAndCodeReq

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: "Invalid request format"})
		return
	}
	defer r.Body.Close()

	err = req.IsOtpIDAndCodeValid()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: err.Error()})
		return
	}
	
	user, session, err := h.AuthMethodSrv.VerifyMfaSmsOtp(context.Background(), req.OtpID, req.OtpCode)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: err.Error()})
		return
	}

	resp := core.NewAuthResp(user, session)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
