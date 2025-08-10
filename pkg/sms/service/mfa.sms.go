package service

import (
	"fmt"
)

func (s *Service) SendMfaSms(countryCode, to, code string) error {
	msg := fmt.Sprintf("Your sign-in verification code is: %s", code)
	completeNumber := fmt.Sprintf("+%s%s", countryCode, to)
	return s.SmsAdp.Send(completeNumber, msg)
}
