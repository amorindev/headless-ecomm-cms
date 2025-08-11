package domain

import (
	"time"
)

type Session struct {
	ID               interface{} `bson:"_id"`
	UserID           interface{} `bson:"user_id" `
	AccessToken      string      `bson:"-"`
	RefreshTokenID   string      `bson:"refresh_token_id"`
	RefreshToken     string      `bson:"-"`
	RefreshTokenHash *string     `bson:"refresh_token"`
	Device           *string     `bson:"device"`
	Revoked          bool        `bson:"revoked"`
	RememberMe       bool        `bson:"remember_me"`
	ExpiresAt        time.Time   `bson:"-"`
	ExpiresIn        int64       `bson:"expires_in"`
	CreatedAt        *time.Time  `bson:"create_at"`
}

func NewSession(userID string, rememberMe bool) *Session {
	return &Session{
		UserID:     userID,
		RememberMe: rememberMe,
	}
}
