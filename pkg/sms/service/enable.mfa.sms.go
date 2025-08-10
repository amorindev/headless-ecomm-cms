package service

import "fmt"

func (s *Service) SendEnableMfaSms(countryCode, to, code string) error {
	msg := fmt.Sprintf("Your verification code to enable SMS two-factor authentication is: %s", code)
	completeNumber := fmt.Sprintf("+%s%s", countryCode,to)
	return s.SmsAdp.Send(completeNumber, msg)
}
