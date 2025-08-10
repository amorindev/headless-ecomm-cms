package domain

import (
	"time"

	"com.fernando/pkg/app/auth-methods/domain"
	//roleModel "com.fernando/pkg/app/role/model"
)

type User struct {
	ID               interface{}                `bson:"_id,omitempty"`
	Email            string                     `bson:"email"`
	EmailVerified    bool                       `bson:"email_verified" db:"email_verified"`
	Name             *string                    `bson:"name"`
	Roles            []string                   `bson:"-"`
	IsActive         bool                       `bson:"is_active"`
	UserPasswordAuth *domain.UserPasswordAuth   `bson:"pass_method"`
	AuthProviders    []*domain.UserProviderAuth `bson:"auth_providers"`
	RolesIDs         []string                   `bson:"roles_ids"`
	CreatedAt        *time.Time                 `bson:"created_at"`
	UpdatedAt        *time.Time                 `bson:"updated_at"`
	MfaStatus        *MfaStatus                 `bson:"mfa_status"`
}

type MfaStatus struct {
	IsMfaSmsCompleted bool `bson:"is_mfa_sms_completed"`
}

func (mfaStatus *MfaStatus) ResetMfaStatus() {
	mfaStatus.IsMfaSmsCompleted = false
}

func NewUserSignUp(email string, name *string, password string) *User {
	now := time.Now().UTC()
	return &User{
		Email:         email,
		EmailVerified: false,
		Name:          name,
		AuthProviders: []*domain.UserProviderAuth{},
		UserPasswordAuth: &domain.UserPasswordAuth{
			Password: password,
		},
		IsActive:  true,
		CreatedAt: &now,
		UpdatedAt: &now,
		MfaStatus: nil,
	}
}
