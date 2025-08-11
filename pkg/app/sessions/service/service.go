package service

import (
	"github.com/amorindev/headless-ecomm-cms/internal/auth"
	"github.com/amorindev/headless-ecomm-cms/pkg/app/sessions/port"
)

var _ port.SessionSrv = &Service{}

type Service struct {
	SessionRepository port.SessionRepo
	TokenSrv         *auth.TokenSrv
}

func NewSessionSrv(
	sessionRepo port.SessionRepo,
	tokenSrv *auth.TokenSrv,
	) *Service {
	return &Service{
		SessionRepository: sessionRepo,
		TokenSrv: tokenSrv,
	}
}
