package sale

import (
	"github.com/go-chi/chi"
	"github.com/mkp-pos-cashier-api/internal/domain/sale/service"
	"github.com/mkp-pos-cashier-api/transport/http/middleware"
)

type SaleHandler struct {
	SaleService    service.SaleService
	Authentication *middleware.Authentication
}

func ProvideSaleHandler(saleService service.SaleService, auth *middleware.Authentication) SaleHandler {
	return SaleHandler{
		SaleService:    saleService,
		Authentication: auth,
	}
}

func (h *SaleHandler) Router(r chi.Router) {

	r.Route("/", func(r chi.Router) {
		r.Use(h.Authentication.VerifyJWT)
		r.Group(func(r chi.Router) {
			r.Use(h.Authentication.IsCashier)
			r.Post("/products", h.CreateNewProduct)
			r.Get("/products", h.ViewProduct)
		})
	})
}
