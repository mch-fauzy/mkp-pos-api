package user

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/mkp-pos-cashier-api/internal/domain/auth/model/dto"
	"github.com/mkp-pos-cashier-api/transport/http/response"
)

// RegisterCashier Register user with cashier role
// @Summary Register cashier
// @Description This endpoint for register an user with cashier role.
// @Tags cashier
// @Param request body dto.RegisterCashierRequest true "Required body to register cashier"
// @Produce json
// @Success 201 {object} response.Base
// @Failure 400 {object} response.Base
// @Failure 404 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v1/cashier/register [post]
func (h *AuthHandler) RegisterCashier(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.WithError(w, err)
		return
	}

	var request dto.RegisterCashierRequest
	err = json.Unmarshal(body, &request)
	if err != nil {
		response.WithError(w, err)
		return
	}

	err = request.Validate()
	if err != nil {
		response.WithError(w, err)
		return
	}

	msg, err := h.AuthService.RegisterCashier(request)
	if err != nil {
		response.WithError(w, err)
		return
	}
	response.WithMessage(w, http.StatusCreated, msg)
}

// LoginCashier Login for cashier role
// @Summary Login cashier
// @Description This endpoint for cashier login.
// @Tags cashier
// @Param request body dto.LoginCashierRequest true "Required body to login cashier"
// @Produce json
// @Success 200 {object} response.Base
// @Failure 400 {object} response.Base
// @Failure 404 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v1/cashier/login [post]
func (h *AuthHandler) LoginCashier(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.WithError(w, err)
		return
	}

	var request dto.LoginCashierRequest
	err = json.Unmarshal(body, &request)
	if err != nil {
		response.WithError(w, err)
		return
	}

	err = request.Validate()
	if err != nil {
		response.WithError(w, err)
		return
	}

	result, err := h.AuthService.LoginCashier(request)
	if err != nil {
		response.WithError(w, err)
		return
	}
	response.WithJSON(w, http.StatusOK, result)
}
