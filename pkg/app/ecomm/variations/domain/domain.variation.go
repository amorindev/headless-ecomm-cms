package domain

import "github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/var-options/domain"

type Variation struct {
	ID      interface{}               `json:"id" bson:"_id"`
	Name    string                    `json:"name" bson:"name"`
	Options []*domain.VariationOption `json:"options,omitempty" bson:"options,omitempty"`
}

func NewVariation(name string) *Variation {
	return &Variation{
		Name: name,
	}
}
