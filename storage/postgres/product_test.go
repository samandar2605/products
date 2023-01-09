package postgres_test

import (
	"testing"

	"github.com/bxcodec/faker/v4"
	"github.com/samandar2605/products/storage/repo"
	"github.com/stretchr/testify/require"
)

func TestGetAllUser(t *testing.T) {
	num, err := faker.RandomInt(100)
	require.NoError(t, err)
	users, err := strg.Product().GetAllProducts(&repo.GetProductsRequest{
		Limit: num[0],
		Page:  1,
	})
	require.NoError(t, err)
	require.NotEmpty(t, users)
}
