package domain

import "time"

type Phone struct {
	ID          interface{} `json:"id" bson:"_id"`
	UserID      interface{} `json:"user_id" bson:"user_id"`
	Number      string      `json:"number" bson:"number"`
	CountryCode string      `json:"country_code" bson:"country_code"`
	CountryIso  string      `json:"country_iso" bson:"country_iso"`
	IsDefault   bool        `json:"is_default" bson:"is_default"`
	IsVerified  bool        `json:"is_verified" bson:"is_verified"`
	CreatedAt   *time.Time  `json:"created_at" bson:"created_at"`
	UpdatedAt   *time.Time  `json:"updated_at" bson:"updated_at"`
}
