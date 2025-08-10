package resend

import (
	"log"
	"os"

	"github.com/resend/resend-go/v2"
)

func NewClient() *resend.Client {
	apiKey := os.Getenv("RESEND_API_KEY")
	if apiKey == "" {
		log.Fatal("RESEND_API_KEY should not be empty")
	}
	return resend.NewClient(apiKey)
}
