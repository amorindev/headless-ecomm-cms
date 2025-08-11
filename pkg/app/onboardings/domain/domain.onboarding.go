package domain

import (
	"time"

	"github.com/amorindev/headless-ecomm-cms/pkg/app/onboardings/core"
)

// Onboarding represents an introductory or informational content
// shown to the user when they first use the system or during
// specific stages. It can include text, a title, and optionally
// an attached file.
type Onboarding struct {
	ID          interface{} `json:"id" bson:"_id"`
	FilePath    string      `json:"file_path,omitempty" bson:"file_name"`
	File        []byte      `json:"-" bson:"-"`
	ContentType string      `json:"content_type,omitempty" bson:"-"`
	ImgUrl      string      `json:"img_url" bson:"-"`
	Title       string      `json:"title" bson:"title"`
	Text        string      `json:"text" bson:"text"`
	Seen        bool        `json:"-" bson:"seen"`
	ExpiresAt   *time.Time  `json:"expires_at" bson:"expires_at"`
	CreatedAt   *time.Time  `json:"created_at" bson:"created_at"`
}

func NewFromCore(oCore *core.OnboardingCore) *Onboarding {
	return &Onboarding{
		Title:       oCore.Title,
		Text:        oCore.Text,
		ContentType: oCore.ContentType,
		FilePath:    oCore.FilePath,
		File:        oCore.File,
		ExpiresAt:   oCore.ExpiresAt,
	}
}
