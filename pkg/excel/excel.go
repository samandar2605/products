package excel

import (
	"strconv"

	"github.com/samandar2605/products/storage/repo"
	"github.com/xuri/excelize/v2"
)

func ReadExcel(filename string) (*[]repo.Product, error) {
	var products []repo.Product
	path := "./media/"
	f, err := excelize.OpenFile(path + filename)

	if err != nil {
		return &[]repo.Product{}, err
	}
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		return &[]repo.Product{}, err
	}

	for i, row := range rows {
		product := repo.Product{}
		if i == 0 {
			continue
		}
		for j, col := range row {
			if j == 0 {
				id, err := strconv.Atoi(col)
				if err != nil {
					return &[]repo.Product{}, err
				}
				product.Id = id
			} else if j == 1 {
				product.Name = col
			} else if j == 2 {
				product.Sku = col
			} else if j == 3 {
				product.Description = col
			} else if j == 4 {
				price, err := strconv.ParseFloat(col, 64)
				if err != nil {
					return &[]repo.Product{}, err
				}
				product.Price = price
			} else if j == 5 {
				count, err := strconv.Atoi(col)
				if err != nil {
					return &[]repo.Product{}, err
				}
				product.Count = count
			}
		}
		products = append(products, product)
	}
	return &products, nil
}
