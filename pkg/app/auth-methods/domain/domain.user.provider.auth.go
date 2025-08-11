package domain

import (
	"time"
)

// UserProviderAuth represents the user's account information from an external authentication provider
// (e.g., Google, Facebook, Apple) when linking or adding a new provider to the user's account.
type UserProviderAuth struct {
	ID            interface{} `json:"-" bson:"_id,omitempty"`
	UserID        interface{} `json:"-" bson:"user_id"`
	Provider      string      `json:"provider" bson:"provider"`
	ProviderID    string      `json:"provider_id" bson:"provider_id"`
	ProviderEmail string      `json:"provider_email"`
	ProviderName  string      `json:"provider_name"`
	AvatarUrl     string      `json:"avatar_url"`
	EmailVerified string      `json:"email_verified"`
	Locale        string      `json:"locale"`
	ExpiresAt     *time.Time  `json:"expires_at"`
	CreatedAt     *time.Time  `json:"created_at" bson:"created_at"`
	UpdatedAt     *time.Time  `json:"updated_at" bson:"updated_at"`
}

