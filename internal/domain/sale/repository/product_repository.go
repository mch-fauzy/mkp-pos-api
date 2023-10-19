package repository

import (
	"fmt"

	"github.com/mkp-pos-cashier-api/internal/domain/sale/model"
	"github.com/mkp-pos-cashier-api/shared/failure"
	"github.com/rs/zerolog/log"
)

const (
	createNewProductQuery = `
	INSERT INTO "product" (name, category, stock, created_at, created_by, updated_at, updated_by) 
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	checkProductByIdQuery = `
	SELECT
		COUNT(id)
	FROM
		"product"
	WHERE
		id = $1
	`
)

type ProductRepository interface {
	CreateProduct(createtProduct *model.CreateProduct) error
}

func (r *SaleRepositoryPostgres) CreateProduct(createtProduct *model.CreateProduct) error {
	query := fmt.Sprintf(createNewProductQuery)
	_, err := r.DB.Write.Exec(
		query,
		createtProduct.Name,
		createtProduct.Category,
		createtProduct.Stock,
		createtProduct.CreatedAt,
		createtProduct.CreatedBy,
		createtProduct.UpdatedAt,
		createtProduct.UpdatedBy,
	)
	if err != nil {
		log.Error().Err(err).Msg("[CreateProduct] Failed exec create product query")
		return err
	}

	return nil
}

func (r *SaleRepositoryPostgres) IsExistProductById(id int) (bool, error) {
	query := fmt.Sprintf(checkProductByIdQuery)
	count := 0
	err := r.DB.Read.Get(&count, query, id)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistProductById] Failed to check product")
		err = failure.InternalError(err)
		return false, err
	}
	return count > 0, nil
}
