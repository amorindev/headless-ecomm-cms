package port

type SmsAdp interface {
	Send(to, msg string) error
}

type SmsSrv interface {
	SendEnableMfaSms(countryCode, to, code string) error
	SendMfaSms(countryCode, to, code string) error
}
