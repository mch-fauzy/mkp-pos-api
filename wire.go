//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/mkp-pos-cashier-api/configs"
	"github.com/mkp-pos-cashier-api/infras"
	userRepository "github.com/mkp-pos-cashier-api/internal/domain/user/repository"
	userService "github.com/mkp-pos-cashier-api/internal/domain/user/service"
	userHandler "github.com/mkp-pos-cashier-api/internal/handlers/user"
	"github.com/mkp-pos-cashier-api/transport/http"
	"github.com/mkp-pos-cashier-api/transport/http/middleware"
	"github.com/mkp-pos-cashier-api/transport/http/router"
)

// Wiring for configurations.
var configurations = wire.NewSet(
	configs.Get,
)

// Wiring for persistences.
var persistences = wire.NewSet(
	infras.ProvidePostgreSQLConn,
)

// Wiring for domain.
var domainUser = wire.NewSet(
	// Service interface and implementation
	userService.ProvideUserServiceImpl,
	wire.Bind(new(userService.UserService), new(*userService.UserServiceImpl)),
	// Repository interface and implementation
	userRepository.ProvideUserRepositoryPostgres,
	wire.Bind(new(userRepository.UserRepository), new(*userRepository.UserRepositoryPostgres)),
)

// Wiring for all domains.
var domains = wire.NewSet(
	domainUser,
)

var authMiddleware = wire.NewSet(
	middleware.ProvideAuthentication,
)

// Wiring for HTTP routing.
var routing = wire.NewSet(
	wire.Struct(new(router.DomainHandlers), "UserHandler"),
	userHandler.ProvideUserHandler,
	router.ProvideRouter,
)

// Wiring for everything.
func InitializeService() *http.HTTP {
	wire.Build(
		// configurations
		configurations,
		// persistences
		persistences,
		// domains
		domains,
		// middleware
		// authMiddleware,
		// routing
		routing,
		// selected transport layer
		http.ProvideHTTP)
	return &http.HTTP{}
}
