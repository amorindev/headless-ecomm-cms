package domain

type VariationOption struct {
	ID          interface{} `bson:"_id" json:"_id"`
	Name        string      `bson:"name" json:"name"`
	Value       string      `bson:"value,omitempty" json:"value"`
	VariationID interface{} `bson:"variation_id" json:"variation_id"`
}

func NewVarOpt(name, value string) *VariationOption {
	return &VariationOption{
		Name:  name,
		Value: value,
	}
}

func NewVariationOption(name, value, variationID string) *VariationOption {
	return &VariationOption{
		Name:        name,
		Value:       value,
		VariationID: variationID,
	}
}
