package service

import (
	"github.com/mkp-pos-cashier-api/internal/domain/sale/model/dto"
	"github.com/rs/zerolog/log"
)

type ProductService interface {
	CreateNewProduct(req dto.CreateProductRequest) (string, error)
}

func (s *SaleServiceImpl) CreateNewProduct(req dto.CreateProductRequest) (string, error) {
	message := "Failed"

	newProduct := req.ToModel()
	err := s.SaleRepository.CreateProduct(&newProduct)
	if err != nil {
		log.Error().Err(err).Msg("[CreateNewProduct] Failed to create new product")
		return message, err
	}

	message = "Success"
	return message, nil
}
