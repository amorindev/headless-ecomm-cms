package core

import "time"

type OnboardingCore struct {
	Title       string     `json:"title"`
	Text        string     `json:"text"`
	ImgUrl      string     `json:"img_url"`
	ContentType string     `json:"content_type"`
	FilePath    string     `json:"file_path"`
	File        []byte     `json:"-"`
	ExpiresAt   *time.Time `json:"expires_at"`
}
