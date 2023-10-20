package service

import (
	"github.com/mkp-pos-cashier-api/internal/domain/product/repository"
)

type ProductService interface {
	ProductManagementService
}

type ProductServiceImpl struct {
	ProductRepository repository.ProductRepository
}

func ProvideProductServiceImpl(productRepository repository.ProductRepository) *ProductServiceImpl {
	return &ProductServiceImpl{
		ProductRepository: productRepository,
	}
}
