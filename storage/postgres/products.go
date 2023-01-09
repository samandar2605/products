package postgres

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/samandar2605/products/storage/repo"
)

type productRepo struct {
	db *sqlx.DB
}

func NewProduct(db *sqlx.DB) repo.ProductStorageI {
	return &productRepo{db: db}
}

func (pr *productRepo) InserterProducts(products *[]repo.Product) error {

	insertQuery := `
		INSERT INTO products(
			name,
			sku,
			description,
			price,
			count
		)values($1,$2,$3,$4,$5)
	`

	updateQuery := `
		UPDATE products SET 
			name=$1,
			description=$2,
			price=$3,
			count=count+$4
		WHERE sku=$5
	`

	isThereQuery := `
		SELECT
			id
		from products
		where sku=$1
	`

	for _, product := range *products {
		tempId := -1
		qator := pr.db.QueryRow(isThereQuery, product.Sku)
		if err := qator.Scan(&tempId); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				continue
			} else {
				return err
			}
		}

		if tempId != -1 {
			_, err := pr.db.Exec(
				updateQuery,
				product.Name,
				product.Description,
				product.Price,
				product.Count,
				product.Sku,
			)
			if err != nil {
				return err
			}
		} else {
			_, err := pr.db.Exec(
				insertQuery,
				product.Name,
				product.Sku,
				product.Description,
				product.Price,
				product.Count,
			)
			if err != nil {
				return err

			}
		}
	}

	return nil
}

// func (pr *productRepo) GetProduct(Id int) (*repo.Product, error)
func (pr *productRepo) GetAllProducts(param *repo.GetProductsRequest) (*repo.GetAllProductsResult, error) {
	result := repo.GetAllProductsResult{
		Products: make([]*repo.Product, 0),
	}

	offset := (param.Page - 1) * param.Limit

	limit := fmt.Sprintf(" LIMIT %d OFFSET %d ", param.Limit, offset)
	filter := "WHERE true"
	if param.Search != "" {
		str := "%" + param.Search + "%"
		filter += fmt.Sprintf(` 
			and name ILIKE '%s' OR sku ILIKE '%s' OR description ILIKE '%s'`, str, str, str)
	}

	if param.SortByDate == "" {
		param.SortByDate = "desc"
	}

	query := `
		SELECT 
			id,
			name,
			sku,
			description,
			price,
			count,
			created_at
		FROM products
		` + filter + `
		ORDER BY created_at ` + param.SortByDate + ` ` + limit

	rows, err := pr.db.Query(query)
	if err != nil {
		return &repo.GetAllProductsResult{}, err
	}

	defer rows.Close()
	for rows.Next() {
		var product repo.Product
		if err := rows.Scan(
			&product.Id,
			&product.Name,
			&product.Sku,
			&product.Description,
			&product.Price,
			&product.Count,
			&product.CreatedAt,
		); err != nil {
			return &repo.GetAllProductsResult{}, err
		}
		result.Products = append(result.Products, &product)
	}

	queryCount := `SELECT count(1) FROM products ` + filter
	err = pr.db.QueryRow(queryCount).Scan(&result.Count)
	if err != nil {
		return &repo.GetAllProductsResult{}, err
	}
	return &result, nil
}

// func (pr *productRepo) UpdateProduct(reqUpdate *repo.Product) (*repo.Product, error)
// func (pr *productRepo) DeleteProduct(Id int) error
