package service

import (
	"context"
	"errors"
	"time"

	userD "github.com/amorindev/headless-ecomm-cms/pkg/app/users/domain"
	userErr "github.com/amorindev/headless-ecomm-cms/pkg/app/users/errors"


)


func (s *Service) SignUp(ctx context.Context, userParam *userD.User) error {
	user, err := s.UserRepo.FindByEmail(ctx, userParam.Email)
	if err != nil {
		if err != userErr.ErrUserNotFound {
			return err
		}
	}

	if user != nil {
		return errors.New("email-already-in-use")
	}

	err = userParam.UserPasswordAuth.HashPassword()
	if err != nil {
		return err
	}

	// * Create the user
	now := time.Now().UTC()
	userParam.UserPasswordAuth.CreatedAt = &now
	userParam.UserPasswordAuth.UpdatedAt = &now
	userParam.IsActive = true
	userParam.EmailVerified = false

	// * assign roles
	userRoles := []string{"User"}
	roleIDs, err := s.RoleRepo.FindIDs(ctx, userRoles)
	if err != nil {
		return err
	}

	userParam.Roles = userRoles

	// * Insert in the DB
	err = s.UserRepo.InsertWithRoles(ctx, userParam, roleIDs)
	if err != nil {
		return err
	}

	userParam.UserPasswordAuth = nil

	return nil
}
