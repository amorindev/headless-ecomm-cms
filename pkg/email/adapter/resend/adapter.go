package resend

import (
	"github.com/amorindev/headless-ecomm-cms/pkg/email/port"
	"github.com/resend/resend-go/v2"
)

var _ port.EmailAdapter = &ResendAdapter{}

type ResendAdapter struct {
	Client *resend.Client
	Form string
}

func NewAdapter(client *resend.Client, from string) *ResendAdapter{
	return &ResendAdapter{
		Client: client,
		Form: from,
	}
}