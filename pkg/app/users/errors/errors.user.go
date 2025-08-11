package errors

import "errors"

var (
	ErrUserNotFound   = errors.New("user not found")
	ErrMfaSmsNotFound = errors.New("mfa sms not found")
)
