package repo

import (
	_ "github.com/golang/mock/mockgen/model"
	"time"
)

type Product struct {
	Id          int       `db:"id"`
	Name        string    `db:"name"`
	Sku         string    `db:"sku"`
	Description string    `db:"description"`
	Price       float64   `db:"price"`
	Count       int       `db:"count"`
	CreatedAt   time.Time `db:"created_at"`
}

type InsertProduct struct {
	Name        string  `db:"name"`
	Sku         string  `db:"sku"`
	Description string  `db:"description"`
	Price       float64 `db:"price"`
	Count       int     `db:"count"`
}

type ProductStorageI interface {
	InserterProducts(*[]Product) error
	// GetProduct(int) (*Product, error)
	GetAllProducts(*GetProductsRequest) (*GetAllProductsResult, error)
	// UpdateProduct(*Product) (*Product, error)
	// DeleteProduct(int) error
}

type GetProductsRequest struct {
	Page       int    `json:"page" db:"page" binding:"required" default:"1"`
	Limit      int    `json:"limit" db:"limit" binding:"required" default:"10"`
	Search     string `json:"search"`
	SortByDate string `json:"sort_by_date" enums:"asc,desc" default:"desc"`
}

type GetAllProductsResult struct {
	Products []*Product
	Count    int
}
