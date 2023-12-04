package user

import (
	"github.com/go-chi/chi"
	"github.com/mkp-pos-cashier-api/internal/domain/auth/service"
)

type AuthHandler struct {
	AuthService service.AuthService
}

func ProvideAuthHandler(authService service.AuthService) AuthHandler {
	return AuthHandler{
		AuthService: authService,
	}
}

func (h *AuthHandler) Router(r chi.Router) {

	r.Route("/cashier", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Post("/register", h.RegisterCashier)
			r.Post("/login", h.LoginCashier)
		})
	})
}
