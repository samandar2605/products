package postgres_test

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/samandar2605/products/storage"
)

var (
	strg storage.StorageI
)

func TestMain(m *testing.M) {

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"localhost",
		"5432",
		"postgres",
		"12345",
		"products",
	)

	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("failed to open connection: %v", err)
	}

	strg = storage.NewStoragePg(db)
	os.Exit(m.Run())
}
