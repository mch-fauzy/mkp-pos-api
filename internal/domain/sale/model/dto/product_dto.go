package dto

import (
	"regexp"
	"strconv"
	"time"

	"github.com/mkp-pos-cashier-api/internal/domain/sale/model"
	"github.com/mkp-pos-cashier-api/shared/failure"
)

type CreateProductRequest struct {
	Name     string `json:"name"`
	Category string `json:"category"`
	Stock    int    `json:"stock"`
	Username string `json:"-"`
}

const (
	notDigitPattern = "^[0-9]+$"
)

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