package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/amorindev/headless-ecomm-cms/cmd/api/message"
	"github.com/amorindev/headless-ecomm-cms/pkg/app/auth-methods/core"
	"github.com/amorindev/headless-ecomm-cms/pkg/app/users/domain"
)

// SignUp handles user registration requests
// Validates the request, creates a new user, and initiates OTP verification
func (h Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req core.SignUpReq

	// Decode JSON request body into SignUpReq struct
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: "Invalid request format"})
		return
	}

	defer r.Body.Close()

	// Validate the sign-up request
	err = req.IsSignUpValid()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: err.Error()})
		return
	}

	// Create new user from validated request data
	// Note: ConfirmPassword is excluded as it's only used for validation
	user := domain.NewUserSignUp(req.Email, req.Name, req.Password)

	err = h.AuthMethodSrv.SignUp(context.TODO(), user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: err.Error()})
		//json.NewEncoder(w).Encode(message.ErrorMessage{Message: "Registration error. Please try again"})
		return
	}

	// Create response using the unified AuthResp structure
	// For sign up: Session is null, OtpID is provided, VerificationRequired is true
	resp := core.NewSignUpResp(user)

	// Return successful response
	// Note: User session will be created after email verification
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
