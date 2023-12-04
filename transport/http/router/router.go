// This generated by evm-cli, edit as necessary
package router

import (
	"github.com/go-chi/chi"
	authHandler "github.com/mkp-pos-cashier-api/internal/handlers/auth"
	productHandler "github.com/mkp-pos-cashier-api/internal/handlers/product"
)

// DomainHandlers is a struct that contains all domain-specific handlers.
type DomainHandlers struct {
	AuthHandler    authHandler.AuthHandler
	ProductHandler productHandler.ProductHandler
}

// Router is the router struct containing handlers.
type Router struct {
	DomainHandlers DomainHandlers
}

// ProvideRouter is the provider function for this router.
func ProvideRouter(domainHandlers DomainHandlers) Router {
	return Router{
		DomainHandlers: domainHandlers,
	}
}

// SetupRoutes sets up all routing for this server.
func (r *Router) SetupRoutes(mux *chi.Mux) {
	mux.Route("/v1", func(rc chi.Router) {
		r.DomainHandlers.AuthHandler.Router(rc)
		r.DomainHandlers.ProductHandler.Router(rc)
	})
}
