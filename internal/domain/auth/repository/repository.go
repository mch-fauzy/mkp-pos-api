package repository

import (
	"github.com/mkp-pos-cashier-api/infras"
)

type AuthRepository interface {
	AuthManagementRepository
}

type AuthRepositoryPostgres struct {
	DB *infras.PostgreSQLConn
}

func ProvideAuthRepositoryPostgres(db *infras.PostgreSQLConn) *AuthRepositoryPostgres {
	return &AuthRepositoryPostgres{
		DB: db,
	}
}
