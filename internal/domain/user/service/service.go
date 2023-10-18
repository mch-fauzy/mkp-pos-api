package service

import (
	"github.com/mkp-pos-cashier-api/configs"
	"github.com/mkp-pos-cashier-api/internal/domain/user/repository"
)

type UserService interface {
	CashierService
}

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	CFG            *configs.Config
}

func ProvideUserServiceImpl(userRepository repository.UserRepository, cfg *configs.Config) *UserServiceImpl {
	return &UserServiceImpl{
		UserRepository: userRepository,
		CFG:            cfg,
	}
}
