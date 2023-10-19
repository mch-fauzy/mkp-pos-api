package repository

import (
	"github.com/mkp-pos-cashier-api/infras"
)

type SaleRepository interface {
	ProductRepository
}

type SaleRepositoryPostgres struct {
	DB *infras.PostgreSQLConn
}

func ProvideSaleRepositoryPostgres(db *infras.PostgreSQLConn) *SaleRepositoryPostgres {
	return &SaleRepositoryPostgres{
		DB: db,
	}
}
