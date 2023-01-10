package storage

import (
	"github.com/jmoiron/sqlx"
	"github.com/samandar2605/products/storage/postgres"
	"github.com/samandar2605/products/storage/repo"
	_ "github.com/golang/mock/mockgen/model"
)

type StorageI interface {
	Product() repo.ProductStorageI
}

type storagePg struct {
	productRepo repo.ProductStorageI
}

func NewStoragePg(db *sqlx.DB) StorageI {
	return &storagePg{
		productRepo: postgres.NewProduct(db),
	}
}

func (s *storagePg) Product() repo.ProductStorageI {
	return s.productRepo
}
