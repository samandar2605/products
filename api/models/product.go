package models

import "time"

type GetProductsRequest struct {
	Page       int    `json:"page" db:"page" binding:"required" default:"1"`
	Limit      int    `json:"limit" db:"limit" binding:"required" default:"10"`
	Search     string `json:"search"`
	SortByDate string `json:"sort_by_date" enums:"asc,desc" default:"desc"`
}

type GetAllProductsResult struct {
	Products []*Product `json:"products"`
	Count    int        `json:"count"`
}

type Product struct {
	Id          int       `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
	Sku         string    `db:"sku" json:"sku"`
	Description string    `db:"description" json:"description"`
	Price       float64   `db:"price" json:"price"`
	Count       int64     `json:"count"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
}
type InsertProduct struct {
	Name        string  `db:"name" json:"name"`
	Sku         string  `db:"sku" json:"sku"`
	Description string  `db:"description" json:"description"`
	Price       float64 `db:"price" json:"price"`
	Count       int64   `json:"count"`
}
