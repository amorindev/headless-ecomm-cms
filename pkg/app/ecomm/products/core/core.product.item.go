package core

type CreateProductReq struct {
	Name             string                `json:"name"`
	Description      string                `json:"description"`
	CategoryID       string                `json:"category_id"`
	Status           string                `json:"status"`
	CreateProductReq *CreateProductItemReq `json:"product"`
}

type CreateProductItemReq struct {
	Stock    int     `json:"stock"`
	Discount int     `json:"discount"`
	Rating   string  `json:"rating"`
	Price    float64 `json:"price"`
}

type CreateProductVariantsReq struct {
	Name                    string                   `json:"name"`
	Description             string                   `json:"description"`
	CategoryID              string                   `json:"category_id"`
	Status                  string                   `json:"status"`
	CreateProductVariantReq *CreateProductVariantReq `json:"variants"`
}

type CreateProductVariantReq struct {
	Stock     int64    `json:"stock"`
	Discount  float32  `json:"discount"`
	Rating    string   `json:"rating"`
	Price     float64  `json:"price"`
	OptionIDS []string `json:"options_ids"`
}
