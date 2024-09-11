package model

import "time"

type ProductPriceDiff struct {
	ID        int       `json:"id"`
	ProductID int       `json:"product_id"`
	OldPrice  float64   `json:"old_price"`
	NewPrice  float64   `json:"new_price"`
	CreatedAt time.Time `json:"created_at"`
	SiteID    int       `json:"site_id"`
}
