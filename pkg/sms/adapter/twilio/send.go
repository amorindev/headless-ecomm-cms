package twilio

import (
	api "github.com/twilio/twilio-go/rest/api/v2010"
)

func (a *Adapter) Send(to, msg string) error {
	params := &api.CreateMessageParams{}
	params.SetBody(msg)
	params.SetFrom(a.FromNumber)
	params.SetTo(to)

	_, err := a.TwilioClient.Api.CreateMessage(params)
	if err != nil {
		return err
	}
	// * Only dev
	/* fmt.Printf("----------------------------------- SE ENVIO EL OTP SMS\n")
	fmt.Printf("----------------------------------- FROM: %s\n", a.FromNumber)
	fmt.Printf("----------------------------------- To: %s\n", to)
	fmt.Printf("----------------------------------- Msg: %s\n", msg) */
	return nil
}
