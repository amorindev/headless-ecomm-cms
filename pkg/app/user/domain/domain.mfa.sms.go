package domain

import "time"

type UserMfaSms struct {
	ID        interface{} `json:"id" bson:"_id"`
	UserID    interface{} `json:"user_id" bson:"user_id"`
	PhoneID   interface{} `json:"phone_id" bson:"phone_id"`
	Confirmed bool        `json:"confirmed" bson:"confirmed"`
	CreatedAt *time.Time  `json:"created_at" bson:"created_at"`
	UpdatedAt *time.Time  `json:"updated_at" bson:"updated_at"`
}

func NewUserMfaSms(userID string, phoneID string, confirmed bool) *UserMfaSms{
	now := time.Now().UTC()
	return &UserMfaSms{
		UserID: userID,
		PhoneID: phoneID,
		Confirmed: confirmed,
		CreatedAt: &now,
		UpdatedAt: &now,
	}
}