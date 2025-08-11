package service

import (
	"context"
	"errors"

	authErr "github.com/amorindev/headless-ecomm-cms/pkg/app/auth-methods/errors"
	sessionD "github.com/amorindev/headless-ecomm-cms/pkg/app/sessions/domain"
	userD "github.com/amorindev/headless-ecomm-cms/pkg/app/users/domain"
	userErr "github.com/amorindev/headless-ecomm-cms/pkg/app/users/errors"
)

func (s *Service) SignIn(ctx context.Context, email string, password string, rememberMe bool) (*userD.User, *sessionD.Session, error) {
	// * Verify if user exists
	user, err := s.UserRepo.FindByEmail(ctx, email)
	if err != nil {
		return nil, nil, err
	}

	if !user.IsActive {
		return nil, nil, errors.New("auth-error")
	}

	_, err = user.UserPasswordAuth.PasswordMatch(password)
	if err != nil {
		if errors.Is(err, authErr.ErrPassDoNotMatch) {
			return nil, nil, errors.New("invalid credentials")
		}
		return nil, nil, err
	}

	roles, err := s.RoleRepo.FindByIDs(ctx, user.RolesIDs)
	if err != nil {
		return nil, nil, err
	}

	user.Roles = roles

	if !user.EmailVerified {
		return user, nil, nil
	}

	mfaSms, err := s.UserRepo.FindMfaSmsByUserID(ctx, user.ID.(string))
	if err != nil {
		if err != userErr.ErrMfaSmsNotFound {
			return nil, nil, err
		}
	}

	if mfaSms != nil && mfaSms.Confirmed {
		return user, nil, nil
	}

	session := sessionD.NewSession(user.ID.(string), rememberMe)

	err = s.SessionSrv.Create(session, roles, email)
	if err != nil {
		return nil, nil, err
	}

	user.UserPasswordAuth = nil

	return user, session, nil
}
