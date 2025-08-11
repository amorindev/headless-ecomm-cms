package port

import (
	"context"

	"github.com/amorindev/headless-ecomm-cms/pkg/app/otp-codes/domain"
)

type OtpRepo interface {
	Insert(ctx context.Context, otp *domain.OtpCodes) error
	Find(ctx context.Context, otpID string) (*domain.OtpCodes, error)
	Delete(ctx context.Context, id string) error
}
