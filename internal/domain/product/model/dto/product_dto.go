package dto

import (
	"regexp"
	"strconv"
	"time"

	"github.com/guregu/null"
	"github.com/mkp-pos-cashier-api/internal/domain/product/model"
	"github.com/mkp-pos-cashier-api/shared"
	"github.com/mkp-pos-cashier-api/shared/failure"
)

const (
	notDigitPattern = "^[0-9]+$"
)

type CreateProductRequest struct {
	Name     string `json:"name"`
	Category string `json:"category"`
	Stock    int    `json:"stock"`
	Username string `json:"-"`
}

func (c CreateProductRequest) Validate() error {
	if c.Name == "" {
		return failure.BadRequestFromString("Name is required")
	}

	if c.Category == "" {
		return failure.BadRequestFromString("Category is required")
	}

	stockStr := strconv.Itoa(c.Stock)
	match, err := regexp.MatchString(notDigitPattern, stockStr)
	if err != nil {
		return failure.InternalError(err)
	}

	if !match {
		return failure.BadRequestFromString("Stock must be a valid number")
	}

	if c.Stock < 0 {
		return failure.BadRequestFromString("Stock must be at least 0")
	}

	return nil
}

func (c CreateProductRequest) ToModel() model.CreateProduct {
	currentTime := time.Now()
	username := c.Username
	return model.CreateProduct{
		Name:      c.Name,
		Category:  c.Category,
		Stock:     c.Stock,
		CreatedAt: currentTime,
		CreatedBy: username,
		UpdatedAt: currentTime,
		UpdatedBy: username,
	}
}

type ViewProductRequest struct {
	Page     int `json:"-"`
	PageSize int `json:"-"`
}

func BuildViewProductRequest(page, pageSize int) ViewProductRequest {
	if page == 0 {
		page = shared.DefaultPage
	}

	if pageSize == 0 {
		pageSize = shared.DefaultPageSize
	}

	return ViewProductRequest{
		Page:     page,
		PageSize: pageSize,
	}
}

func (v ViewProductRequest) Validate() error {
	if v.Page < 0 {
		return failure.BadRequestFromString("Page must be a positive integer")
	}

	if v.PageSize < 0 {
		return failure.BadRequestFromString("PageSize must be a positive integer")
	}

	return nil
}

func (v ViewProductRequest) ToPaginationModel() model.Pagination {
	return model.Pagination{
		Page:     v.Page,
		PageSize: v.PageSize,
	}
}

type ProductResponse struct {
	Id        int         `json:"id"`
	Name      string      `json:"name"`
	Category  string      `json:"category"`
	Stock     int         `json:"stock"`
	CreatedAt time.Time   `json:"createdAt"`
	CreatedBy string      `json:"createdBy"`
	UpdatedAt time.Time   `json:"updatedAt"`
	UpdatedBy string      `json:"updatedBy"`
	DeletedAt null.Time   `json:"deletedAt"`
	DeletedBy null.String `json:"deletedBy"`
}

type ProductListResponse []ProductResponse

func NewProductResponse(product model.Product) ProductResponse {
	return ProductResponse{
		Id:        product.Id,
		Name:      product.Name,
		Category:  product.Category,
		Stock:     product.Stock,
		CreatedAt: product.CreatedAt,
		CreatedBy: product.CreatedBy,
		UpdatedAt: product.UpdatedAt,
		UpdatedBy: product.UpdatedBy,
		DeletedAt: product.DeletedAt,
		DeletedBy: product.DeletedBy,
	}
}

func BuildProductListResponse(productList model.ProductList) ProductListResponse {
	results := ProductListResponse{}
	for _, product := range productList {
		results = append(results, NewProductResponse(*product))
	}
	return results
}
