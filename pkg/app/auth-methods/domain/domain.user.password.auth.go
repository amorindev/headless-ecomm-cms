package domain

import (
	"errors"
	"time"

    authErr "github.com/amorindev/headless-ecomm-cms/pkg/app/auth-methods/errors"
	"golang.org/x/crypto/bcrypt"
)

type UserPasswordAuth struct {
	Password     string     `bson:"-"`
	PasswordHash *string    `bson:"password_hash"`
	CreatedAt    *time.Time `bson:"created_at"`
	UpdatedAt    *time.Time `bson:"updated_at"`
}

func (a *UserPasswordAuth) HashPassword() error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(a.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	pH := string(passwordHash)

	a.PasswordHash = &pH

	return nil
}

func (a UserPasswordAuth) PasswordMatch(password string) (bool, error) {
	if a.PasswordHash == nil {
		return false, errors.New("password hash is nil")
	}
	err := bcrypt.CompareHashAndPassword([]byte(*a.PasswordHash), []byte(password))
	if err != nil {
		return false, authErr.ErrPassDoNotMatch
	}

	return true, nil
}

func NewUserProviderAuth(password string) *UserPasswordAuth {
	return &UserPasswordAuth{
		Password: password,
	}
}
