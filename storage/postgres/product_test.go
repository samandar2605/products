package postgres_test

import (
	"testing"
	"time"

	"github.com/bxcodec/faker/v4"
	"github.com/samandar2605/products/storage/repo"
	"github.com/stretchr/testify/require"
)

func createProduct(t *testing.T) {
	products := []repo.Product{}
	num, err := faker.RandomInt(100)
	require.NoError(t, err)
	products = append(products, repo.Product{
		Id:          num[0],
		Name:        faker.Name(),
		Sku:         faker.UUIDHyphenated(),
		Description: faker.Sentence(),
		Price:       float64(num[0]),
		Count:       num[0],
		CreatedAt:   time.Now(),
	})
	err = strg.Product().InserterProducts(&products)
	require.NoError(t, err)

}

func TestGetProduct(t *testing.T) {
	num, err := faker.RandomInt(100)
	createProduct(t)
	require.NoError(t, err)
	u, err := strg.Product().GetAllProducts(&repo.GetProductsRequest{
		Limit: num[0],
		Page:  1,
	})
	require.NoError(t, err)
	require.NotEmpty(t, u)
}
