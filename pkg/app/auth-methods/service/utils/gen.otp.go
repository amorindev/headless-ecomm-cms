package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// GenOtpCode generates a 6-digit one-time password (OTP) as a zero-padded string.
// The code is cryptographically secure, randomly generated in the range 000000–999999.
func GenOtpCode() (string, error) {
	nBig, err := rand.Int(rand.Reader, big.NewInt(1000000))
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%06d", nBig.Int64()), nil
}
