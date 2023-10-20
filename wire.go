//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/mkp-pos-cashier-api/configs"
	"github.com/mkp-pos-cashier-api/infras"
	productRepository "github.com/mkp-pos-cashier-api/internal/domain/product/repository"
	productService "github.com/mkp-pos-cashier-api/internal/domain/product/service"
	userRepository "github.com/mkp-pos-cashier-api/internal/domain/user/repository"
	userService "github.com/mkp-pos-cashier-api/internal/domain/user/service"
	productHandler "github.com/mkp-pos-cashier-api/internal/handlers/product"
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

var domainProduct = wire.NewSet(
	// Service interface and implementation
	productService.ProvideProductServiceImpl,
	wire.Bind(new(productService.ProductService), new(*productService.ProductServiceImpl)),
	// Repository interface and implementation
	productRepository.ProvideProductRepositoryPostgres,
	wire.Bind(new(productRepository.ProductRepository), new(*productRepository.ProductRepositoryPostgres)),
)

// Wiring for all domains.
var domains = wire.NewSet(
	domainUser,
	domainProduct,
)

var authMiddleware = wire.NewSet(
	middleware.ProvideAuthentication,
)

// Wiring for HTTP routing.
var routing = wire.NewSet(
	wire.Struct(new(router.DomainHandlers), "UserHandler", "ProductHandler"),
	userHandler.ProvideUserHandler,
	productHandler.ProvideProductHandler,
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
		authMiddleware,
		// routing
		routing,
		// selected transport layer
		http.ProvideHTTP)
	return &http.HTTP{}
}
