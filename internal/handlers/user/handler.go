package user

import (
	"github.com/go-chi/chi"
	"github.com/mkp-pos-cashier-api/internal/domain/user/service"
)

type UserHandler struct {
	UserService service.UserService
}

func ProvideUserHandler(userService service.UserService) UserHandler {
	return UserHandler{
		UserService: userService,
	}
}

func (h *UserHandler) Router(r chi.Router) {

	r.Route("/cashier", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Post("/register", h.RegisterCashier)
			r.Post("/login", h.LoginCashier)
		})
	})
}
