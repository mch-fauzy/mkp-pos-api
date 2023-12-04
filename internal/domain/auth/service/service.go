package service

import (
	"github.com/mkp-pos-cashier-api/configs"
	"github.com/mkp-pos-cashier-api/internal/domain/auth/repository"
)

type AuthService interface {
	CashierService
}

type AuthServiceImpl struct {
	AuthRepository repository.AuthRepository
	CFG            *configs.Config
}

func ProvideAuthServiceImpl(authRepository repository.AuthRepository, cfg *configs.Config) *AuthServiceImpl {
	return &AuthServiceImpl{
		AuthRepository: authRepository,
		CFG:            cfg,
	}
}
