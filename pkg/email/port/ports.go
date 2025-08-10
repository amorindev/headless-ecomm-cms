package port

type EmailAdapter interface {
	Send(to, subject, htmlBody string) error
}

type EmailSrv interface {
	SendVerifyEmailOtp(email string, code string) error
	SendVerifyEnableMfaSmsOtp(email string, code string) error
	SendVerifyMfaSmsOtp(email string, code string) error
}
