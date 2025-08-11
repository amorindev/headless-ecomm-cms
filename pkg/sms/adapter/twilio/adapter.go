package twilio

import (
	"github.com/amorindev/headless-ecomm-cms/pkg/sms/port"
	"github.com/twilio/twilio-go"
)

var _ port.SmsAdp = &Adapter{}

type Adapter struct {
	TwilioClient *twilio.RestClient
	FromNumber   string
}

func NewTwilioAdapter(client *twilio.RestClient, from string) *Adapter {
	return &Adapter{
		TwilioClient: client,
		FromNumber:   from,
	}
}
