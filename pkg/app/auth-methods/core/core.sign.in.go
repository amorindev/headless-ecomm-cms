package core

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
)

// SignInReq represents the sign in request structure
type SignInReq struct {
	Email      string `json:"email" validate:"required,email"`
	Password   string `json:"password" validate:"required"`
	RememberMe bool   `json:"remember_me"`
	Platform   string `json:"platform"` // ios, android, web
}

// IsSignInValid performs validation on the sign in request
// Returns an error if any validation fails
func (req SignInReq) IsSignInValid() error {
	// Validate email field is not empty
	if strings.TrimSpace(req.Email) == "" {
		return errors.New("email is required")
	}

	// Validate email format using validator
	validate := validator.New()
	err := validate.Var(req.Email, "email")
	if err != nil {
		return errors.New("invalid email format")
	}

	// Validate password field is not empty
	if strings.TrimSpace(req.Password) == "" {
		return errors.New("password is required")
	}

	return nil
}
