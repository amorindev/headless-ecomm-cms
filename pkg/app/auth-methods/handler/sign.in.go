package handler

import (
	"encoding/json"
	"net/http"

	"github.com/amorindev/headless-ecomm-cms/cmd/api/message"
	"github.com/amorindev/headless-ecomm-cms/pkg/app/auth-methods/core"
)

// SignIn handles user authentication requests
// Validates credentials and returns appropriate response based on user state
func (h Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req core.SignInReq

	// Decode JSON request body into SignInReq struct
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: "Invalid request format"})
		return
	}

	defer r.Body.Close()

	// Validate the sign in request
	err = req.IsSignInValid()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: err.Error()})
		return
	}

	// Authenticate user and get session based on user state
	user, session, err := h.AuthMethodSrv.SignIn(r.Context(), req.Email, req.Password, req.RememberMe)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: err.Error()})
		return
	}

	// Create response
	resp := core.NewSignInResp(user, session)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
