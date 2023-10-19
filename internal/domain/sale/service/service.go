package service

import (
	"github.com/mkp-pos-cashier-api/internal/domain/sale/repository"
)

type SaleService interface {
	ProductService
}

type SaleServiceImpl struct {
	SaleRepository repository.SaleRepository
}

func ProvideSaleServiceImpl(saleRepository repository.SaleRepository) *SaleServiceImpl {
	return &SaleServiceImpl{
		SaleRepository: saleRepository,
	}
}
