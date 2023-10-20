package product

import (
	"github.com/go-chi/chi"
	"github.com/mkp-pos-cashier-api/internal/domain/product/service"
	"github.com/mkp-pos-cashier-api/transport/http/middleware"
)

type ProductHandler struct {
	ProductService service.ProductService
	Authentication *middleware.Authentication
}

func ProvideProductHandler(productService service.ProductService, auth *middleware.Authentication) ProductHandler {
	return ProductHandler{
		ProductService: productService,
		Authentication: auth,
	}
}

func (h *ProductHandler) Router(r chi.Router) {

	r.Route("/", func(r chi.Router) {
		r.Use(h.Authentication.VerifyJWT)
		r.Group(func(r chi.Router) {
			r.Use(h.Authentication.IsCashier)
			r.Post("/products", h.CreateNewProduct)
			r.Get("/products", h.ViewProduct)
		})
	})
}
