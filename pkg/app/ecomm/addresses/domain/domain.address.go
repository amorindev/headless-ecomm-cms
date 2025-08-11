package domain

import "time"

// Address represents a user's or store's saved location details.
type Address struct {
	ID          interface{} `json:"id" bson:"_id"`
	UserID      interface{} `json:"user_id" bson:"user_id"`
	StoreID     interface{} `json:"store_id" bson:"store_id"`
	Label       *string     `json:"label" bson:"label"`
	AddressLine string      `bson:"address_line" json:"address_line"`
	City        string      `bson:"city,omitempty" json:"city,omitempty"`
	State       string      `bson:"state,omitempty" json:"state,omitempty"`
	Country     string      `bson:"country" json:"country"`
	PostalCode  string      `bson:"postal_code" json:"postal_code"`
	Latitude    float64     `bson:"latitude" json:"latitude"`
	Longitude   float64     `bson:"longitude" json:"longitude"`
	IsDefault   bool        `bson:"is_default" json:"is_default"`
	CreatedAt   *time.Time  `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt   *time.Time  `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}

