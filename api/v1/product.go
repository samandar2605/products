package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/samandar2605/products/api/models"
	"github.com/samandar2605/products/storage/repo"
)

func parseProductModel(product *repo.Product) models.Product {
	return models.Product{
		Id:          product.Id,
		Name:        product.Name,
		Sku:         product.Sku,
		Description: product.Description,
		Price:       product.Price,
		Count:       int64(product.Count),
		CreatedAt:   product.CreatedAt,
	}
}

// @Router /products [get]
// @Summary Get all Products
// @Description Get all Products
// @Tags products
// @Accept json
// @Produce json
// @Param filter query models.GetProductsRequest false "Filter"
// @Success 200 {object} models.GetAllProductsResult
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) GetAllProducts(c *gin.Context) {
	req, err := productsParams(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}
	result, err := h.storage.Product().GetAllProducts(&repo.GetProductsRequest{
		Page:       req.Page,
		Limit:      req.Limit,
		Search:     req.Search,
		SortByDate: req.SortByDate,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, productsResponse(result))
}

func productsParams(c *gin.Context) (*models.GetProductsRequest, error) {
	var (
		limit      int = 10
		page       int = 1
		sortByDate string
		err        error
	)

	if c.Query("limit") != "" {
		limit, err = strconv.Atoi(c.Query("limit"))
		if err != nil {
			return nil, err
		}
	}

	if c.Query("page") != "" {
		page, err = strconv.Atoi(c.Query("page"))
		if err != nil {
			return nil, err
		}
	}

	if c.Query("sort_by_date") != "" &&
		(c.Query("sort_by_date") == "desc" || c.Query("sort_by_date") == "asc" || c.Query("sort_by_date") == "none") {
		sortByDate = c.Query("sort_by_date")
	}

	return &models.GetProductsRequest{
		Page:       page,
		Limit:      limit,
		Search:     c.Query("search"),
		SortByDate: sortByDate,
	}, nil
}
func productsResponse(data *repo.GetAllProductsResult) *models.GetAllProductsResult {
	response := models.GetAllProductsResult{
		Products: make([]*models.Product, 0),
		Count:    data.Count,
	}

	for _, product := range data.Products {
		p := parseProductModel(product)
		response.Products = append(response.Products, &p)
	}

	return &response
}
