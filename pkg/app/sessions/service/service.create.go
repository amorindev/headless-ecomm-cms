package service

import (
	"context"
	"time"

	"github.com/amorindev/headless-ecomm-cms/pkg/app/sessions/domain"
)

func (s *Service) Create(session *domain.Session, roles []string, email string) error {
	now := time.Now().UTC()
	session.CreatedAt = &now

	// * AccessToken
	claim, accessToken, err := s.TokenSrv.NewAccessToken(session.UserID.(string), email, roles)
	if err != nil {
		return err
	}

	// * RefreshToken
	rClaim, refreshToken, err := s.TokenSrv.NewRefreshToken(session.UserID.(string), session.RememberMe)
	if err != nil {
		return err
	}

	session.RefreshTokenID = rClaim.ID
	session.RefreshToken = refreshToken
	session.AccessToken = accessToken
	session.ExpiresAt = rClaim.ExpiresAt.Time
	session.ExpiresIn = int64(time.Until(claim.ExpiresAt.Time).Seconds())
	session.Revoked = false
	err = s.SessionRepository.Insert(context.Background(), session)
	if err != nil {
		return err
	}

	return nil
}
