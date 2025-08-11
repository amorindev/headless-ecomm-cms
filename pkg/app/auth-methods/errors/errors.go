package errors

import "errors"

// * Repository layer
var (
	ErrPassMethodNotFound     = errors.New("pass method not found")
	ErrProviderMethodNotFound = errors.New("provider method not found")
)


// * Service layer
var (
	ErrPassDoNotMatch = errors.New("passwords do not match")
	ErrInvalidCredentials = errors.New("invalid credentials") 
)

