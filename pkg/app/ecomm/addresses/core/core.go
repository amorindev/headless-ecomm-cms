package core

import (
	"errors"
	"time"
)

// CreateAddressReq represents the request payload for adding a new address to a user's account.
// It includes location details, optional metadata, and a flag to mark the address as default.
type CreateAddressReq struct {
	Label       *string    `json:"label" `
	AddressLine string     `json:"address_line"`
	Latitude    float64    `json:"latitude"`
	Longitude   float64    `json:"longitude"`
	City        string     `json:"city,omitempty"`
	State       string     `json:"state,omitempty"`
	Country     string     `json:"country"`
	PostalCode  string     `json:"postal_code"`
	IsDefault   bool       `json:"is_default"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}

func (c *CreateAddressReq) IsCreateAddressValid() error {
	if c.PostalCode == "" {
		return errors.New("postal_code is required")
	}
	if c.Label == nil {
		return errors.New("label is required")
	}
	if c.AddressLine == "" {
		return errors.New("address_line is required")
	}
	if c.Country == "" {
		return errors.New("country is required")
	}
	return nil
}
