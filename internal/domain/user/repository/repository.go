package repository

import (
	"github.com/mkp-pos-cashier-api/infras"
)

type UserRepository interface {
	CashierRepository
}

type UserRepositoryPostgres struct {
	DB *infras.PostgreSQLConn
}

func ProvideUserRepositoryPostgres(db *infras.PostgreSQLConn) *UserRepositoryPostgres {
	return &UserRepositoryPostgres{
		DB: db,
	}
}
