package core

import (
	userD "github.com/amorindev/headless-ecomm-cms/pkg/app/users/domain"
	userC "github.com/amorindev/headless-ecomm-cms/pkg/app/users/core"
	sessionD "github.com/amorindev/headless-ecomm-cms/pkg/app/sessions/domain"
	sessionC "github.com/amorindev/headless-ecomm-cms/pkg/app/sessions/core"
	authMethodD "github.com/amorindev/headless-ecomm-cms/pkg/app/auth-methods/domain"
)

// AuthResp represents the unified authentication response structure
// This structure handles all authentication scenarios (sign up, sign in, verification, etc.)
// Fields are conditionally populated based on the authentication flow and user state
type AuthResp struct {
	// Session contains authentication tokens (null for sign up before email verification)
	Session *sessionC.SessionCore `json:"session"`

	// User contains user profile information (always present)
	User *userC.UserCore `json:"user"`

	// Providers contains authentication provider information (null if no providers)
	Providers []*authMethodD.UserProviderAuth `json:"providers"`

	// MfaStatus contains multi-factor authentication status (null if MFA not enabled)
	MfaStatus *MfaStatusCore `json:"mfa_status"`
}

// MfaStatusCore represents multi-factor authentication status
type MfaStatusCore struct {
	IsMfaSmsCompleted bool `json:"is_mfa_sms_completed"`
}

// NewAuthResp creates a new AuthResp with flexible field population
// This function handles all authentication scenarios
func NewAuthResp(
	user *userD.User,
	session *sessionD.Session,
) *AuthResp {
	resp := &AuthResp{
		User: &userC.UserCore{
			ID:            user.ID.(string),
			Email:         user.Email,
			EmailVerified: user.EmailVerified,
			Name:          user.Name,
			Roles:         user.Roles,
			CreatedAt:     user.CreatedAt,
			UpdatedAt:     user.UpdatedAt,
		},
		Providers: nil,
	}

	// Only include session if provided (null for sign up before verification)
	if session != nil {
		resp.Session = &sessionC.SessionCore{
			AccessToken:  session.AccessToken,
			RefreshToken: session.RefreshToken,
			ExpiresIn:    session.ExpiresIn,
		}
	}

	if user.MfaStatus != nil {
		resp.MfaStatus = &MfaStatusCore{
			IsMfaSmsCompleted: user.MfaStatus.IsMfaSmsCompleted,
		}
	}

	return resp
}

// NewSignUpResp creates response for user registration
// Session is null, OtpID is provided for email verification
func NewSignUpResp(user *userD.User) *AuthResp {
	return NewAuthResp(user, nil)
}

// NewSignInResp creates response for successful sign in
// Session is included, OtpID is null unless MFA is required
func NewSignInResp(user *userD.User, session *sessionD.Session) *AuthResp {
	return NewAuthResp(user, session)
}
