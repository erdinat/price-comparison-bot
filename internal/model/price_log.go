package model

import "time"

type ProductPriceDiffLog struct {
	ID        int       `json:"id"`
	ProductID int       `json:"product_id"`
	Data      string    `json:"data"`
	CreatedAt time.Time `json:"created_at"`
}
