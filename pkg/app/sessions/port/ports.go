package port

import (
	"context"

	"github.com/amorindev/headless-ecomm-cms/pkg/app/sessions/domain"
)

type SessionRepo interface {
	Insert(ctx context.Context, session *domain.Session) error
	DeleteByRTokenID(ctx context.Context, rTokenID string) error
}

type SessionSrv interface {
	Create(session *domain.Session, roles []string, email string) error
}
