package service

import (
	"github.com/mkp-pos-cashier-api/internal/domain/sale/model/dto"
	"github.com/rs/zerolog/log"
)

type ProductService interface {
	CreateNewProduct(req dto.CreateProductRequest) (string, error)
	GetProductList(req dto.ViewProductRequest) (dto.ProductListResponse, dto.Pagination, error)
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

func (s *SaleServiceImpl) GetProductList(req dto.ViewProductRequest) (dto.ProductListResponse, dto.Pagination, error) {

	paginationFilter := req.ToPaginationModel()
	product, err := s.SaleRepository.GetProducts(paginationFilter)
	if err != nil {
		log.Error().Err(err).Msg("[GetProductList] Failed to retrieve product list")
		return dto.ProductListResponse{}, dto.Pagination{}, err
	}

	result := dto.BuildProductListResponse(product)
	paginationMetadata := dto.BuildMetadata(req.Page, req.PageSize)
	return result, paginationMetadata, nil
}
