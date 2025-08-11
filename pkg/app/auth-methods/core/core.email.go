package core

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
)

type EmailReq struct {
	Email string `json:"email"`
}

func (req *EmailReq) IsEmailValid() error {
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
	return nil
}
