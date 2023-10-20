package service

import (
	"github.com/mkp-pos-cashier-api/internal/domain/product/model/dto"
	"github.com/rs/zerolog/log"
)

type ProductManagementService interface {
	CreateNewProduct(req dto.CreateProductRequest) (string, error)
	GetProductList(req dto.ViewProductRequest) (dto.ProductListResponse, dto.Pagination, error)
}

func (s *ProductServiceImpl) CreateNewProduct(req dto.CreateProductRequest) (string, error) {
	message := "Failed"

	newProduct := req.ToModel()
	err := s.ProductRepository.CreateProduct(&newProduct)
	if err != nil {
		log.Error().Err(err).Msg("[CreateNewProduct] Failed to create new product")
		return message, err
	}

	message = "Success"
	return message, nil
}

func (s *ProductServiceImpl) GetProductList(req dto.ViewProductRequest) (dto.ProductListResponse, dto.Pagination, error) {

	paginationFilter := req.ToPaginationModel()
	product, err := s.ProductRepository.GetProducts(paginationFilter)
	if err != nil {
		log.Error().Err(err).Msg("[GetProductList] Failed to retrieve product list")
		return dto.ProductListResponse{}, dto.Pagination{}, err
	}

	result := dto.BuildProductListResponse(product)
	paginationMetadata := dto.BuildMetadata(req.Page, req.PageSize)
	return result, paginationMetadata, nil
}
