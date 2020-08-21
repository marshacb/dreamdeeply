package models

import "time"

// Product is something we sell
type Product struct {
	ID          string    `db:"product_id" json:"id"`
	Name        string    `db:"name" json:"name"`
	Price       float64   `db:"price" json:"price"`
	Quantity    int       `db:"quantity" json:"quantity"`
	Sold        int       `db:"sold" json:"sold"`
	Revenue     int       `db:"revenue" json:"revenue"`
	DateCreated time.Time `db:"date_created" json:"date_created"`
	DateUpdated time.Time `db:"date_updated" json:"date_updated"`
}

// NewProduct is what we require from clients when adding a Product.
type NewProduct struct {
	Name     string  `json:"name" validate:"required"`
	Price    float64 `json:"price" validate:"gte=0"`
	Quantity int     `json:"quantity" validate:"gte=1"`
}

// UpdateProduct defines what information may be provided to modify an
// existing Product. All fields are optional so clients can send just the
// fields they want changed. It uses pointer fields so we can differentiate
// between a field that was not provided and a field that was provided as
// explicitly blank. Normally we do not want to use pointers to basic types but
// we make exceptions around marshalling/unmarshalling.
type UpdateProduct struct {
	Name     *string `json:"name"`
	Price    *int    `json:"price" validate:"omitempty,gte=0"`
	Quantity *int    `json:"quantity" validate:"omitempty,gte=1"`
}
