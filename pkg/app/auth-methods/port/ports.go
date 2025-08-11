package port

import (
	"context"

	userD "github.com/amorindev/headless-ecomm-cms/pkg/app/users/domain"
	sessionD "github.com/amorindev/headless-ecomm-cms/pkg/app/sessions/domain"
)

type AuthMethodSrv interface {
	EnableMfaSms(ctx context.Context, userID string, phoneID string) (string, error)

	ResendVerifyEmailOtp(ctx context.Context, email string) (string, error)

	ResendVerifyEnableMfaSmsOtp(ctx context.Context, userID string) (string,error)

	ResendVerifyMfaSmsOtp(ctx context.Context, userID string) (string,error)
    
	SignIn(ctx context.Context, email string, password string, rememberMe bool) (*userD.User, *sessionD.Session, error)

	SignOut(ctx context.Context, rTokenID string) error

	SignUp(ctx context.Context, userParam *userD.User) ( error)

	VerifyEmailOtp(ctx context.Context, otpID string, otpCode string, userID string) (*userD.User, *sessionD.Session, error)

	VerifyEnableMfaSmsOtp(ctx context.Context, otpID string, otpCode string, userID string) error

	VerifyMfaSmsOtp(ctx context.Context, otpID string, otpCode string) (*userD.User, *sessionD.Session, error)
}
