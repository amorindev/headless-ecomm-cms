package twilio

import (
	"github.com/twilio/twilio-go"
)

func NewClient() *twilio.RestClient {
	return twilio.NewRestClient()
}
