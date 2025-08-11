package domain

import (
	"time"
)

type Product struct {
	ID           interface{}    `json:"id" bson:"_id"`
	CategoryID   interface{}    `json:"category_id,omitempty" bson:"category_id"`
	Name         string         `json:"name" bson:"name"`
	Description  string         `json:"description" bson:"description"`
	FilePath     string         `json:"file_path,omitempty" bson:"file_path"`
	File         []byte         `json:"-" bson:"-"`
	ContentType  string         `json:"content_type,omitempty" bson:"-"`
	ImgUrl       string         `json:"img_url" bson:"-"`
	CategoryName string         `json:"category_name,omitempty" bson:"-"`
	CreatedAt    *time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt    *time.Time     `json:"updated_at" bson:"updated_at"`
	ProductItems []*ProductItem `json:"product_items" bson:"product_items,omitempty"`
	Variations   []*Variation   `json:"variations" bson:"-"`
}

// Only response
type Variation struct {
	Name   string   `json:"name"`
	Values []string `json:"values"`
}
