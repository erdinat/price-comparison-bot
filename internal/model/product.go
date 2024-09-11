package model

import "time"

type Product struct {
	ID           int       `json:"id"`
	ProductName  string    `json:"product_name"`
	SKU          string    `json:"sku"`
	Barcode      string    `json:"barcode"`
	CreatedAt    time.Time `json:"created_at"`
	Status       int       `json:"status"`
	ProductImage string    `json:"product_image"`
}
