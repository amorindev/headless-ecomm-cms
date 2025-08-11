package domain

import (
	"time"
)

// ProductItem represents a specific sellable unit of a product, including stock, price, and variation details.
type ProductItem struct {
	ID           interface{}   `json:"id" bson:"_id"`
	Status       string        `json:"status" bson:"status"`
	Sku          string        `json:"sku" bson:"sku"`
	QtyInStock   int           `json:"qty_in_stock" bson:"stock"`
	Discount     int           `json:"discount" bson:"discount"`
	Price        float64       `json:"price" bson:"price"`
	FilePath     string        `json:"file_path,omitempty" bson:"file_path"`
	File         []byte        `json:"-" bson:"-"`
	ContentType  string        `json:"content_type,omitempty" bson:"-"`
	ImgUrl       string        `json:"img_url" bson:"-"`
	CreatedAt    *time.Time    `json:"created_at" bson:"created_at"`
	UpdatedAt    *time.Time    `json:"updated_at" bson:"updated_at"`
	ProductID    interface{}   `json:"-" bson:"-"`
	VarOptionIDs []interface{} `json:"var_option_ids,omitempty" bson:"var_option_ids"`
	Options      []*Option     `json:"options" bson:"options,omitempty"`
}

// Option represents a specific variation choice for a product item (e.g., color, size).
type Option struct {
	Name        string `json:"name" bson:"name"`
	VarOptName  string `json:"var_opt_name" bson:"var_opt_name"` // * ese es RED, YELLOW, M, L
	VarOptValue string `json:"var_opt_value" bson:"var_opt_value"`
}
