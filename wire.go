//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/mkp-pos-cashier-api/configs"
	"github.com/mkp-pos-cashier-api/infras"
	authRepository "github.com/mkp-pos-cashier-api/internal/domain/auth/repository"
	authService "github.com/mkp-pos-cashier-api/internal/domain/auth/service"
	productRepository "github.com/mkp-pos-cashier-api/internal/domain/product/repository"
	productService "github.com/mkp-pos-cashier-api/internal/domain/product/service"
	authHandler "github.com/mkp-pos-cashier-api/internal/handlers/auth"
	productHandler "github.com/mkp-pos-cashier-api/internal/handlers/product"
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
var domainAuth = wire.NewSet(
	// Service interface and implementation
	authService.ProvideAuthServiceImpl,
	wire.Bind(new(authService.AuthService), new(*authService.AuthServiceImpl)),
	// Repository interface and implementation
	authRepository.ProvideAuthRepositoryPostgres,
	wire.Bind(new(authRepository.AuthRepository), new(*authRepository.AuthRepositoryPostgres)),
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
	domainAuth,
	domainProduct,
)

var authMiddleware = wire.NewSet(
	middleware.ProvideAuthentication,
)

// Wiring for HTTP routing.
var routing = wire.NewSet(
	wire.Struct(new(router.DomainHandlers), "AuthHandler", "ProductHandler"),
	authHandler.ProvideAuthHandler,
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
