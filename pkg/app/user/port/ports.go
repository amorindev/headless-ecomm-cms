package port

import (
	"context"

	userD "github.com/amorindev/headless-ecomm-cms/pkg/app/user/domain"
)

type UserSrv interface {
	GetUser(ctx context.Context, userID string) (*userD.User, error)
}

type UserRepo interface {
	// * User
	Find(ctx context.Context, id string) (*userD.User, error)
	FindByEmail(ctx context.Context, email string) (*userD.User, error)
	InsertWithRoles(ctx context.Context, user *userD.User, rolesIDs []string) error
	ConfirmEmail(ctx context.Context, userID string) error

	// * MfaSms
	InsertMfaSms(ctx context.Context, twoFaSms *userD.UserMfaSms) error
	ConfirmMfaSms(ctx context.Context, userID string) error
	FindMfaSmsByUserID(ctx context.Context, userID string) (*userD.UserMfaSms, error)

	// * MfaStatus
	UpdateMfaStatus(ctx context.Context, userID string, mfaStatus *userD.MfaStatus) error
	ResetMfaStatus(ctx context.Context, userID string, mfaStatus *userD.MfaStatus) error
}

