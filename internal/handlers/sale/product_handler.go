package sale

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/mkp-pos-cashier-api/internal/domain/sale/model/dto"
	"github.com/mkp-pos-cashier-api/shared"
	"github.com/mkp-pos-cashier-api/transport/http/response"
)

// CreateNewProduct Create a new product
// @Summary Create a new product
// @Description This endpoint for create a new product.
// @Tags sale
// @Param request body dto.CreateProductRequest true "Required body to create a new product"
// @Produce json
// @Success 201 {object} response.Base
// @Failure 400 {object} response.Base
// @Failure 404 {object} response.Base
// @Failure 500 {object} response.Base
// @Security BearerAuth
// @Router /v1/products [post]
func (h *SaleHandler) CreateNewProduct(w http.ResponseWriter, r *http.Request) {
	username, err := shared.GetUsernameFromContext(r)
	if err != nil {
		response.WithError(w, err)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.WithError(w, err)
		return
	}

	var request dto.CreateProductRequest
	err = json.Unmarshal(body, &request)
	if err != nil {
		response.WithError(w, err)
		return
	}

	request.Username = username
	err = request.Validate()
	if err != nil {
		response.WithError(w, err)
		return
	}

	msg, err := h.SaleService.CreateNewProduct(request)
	if err != nil {
		response.WithError(w, err)
		return
	}
	response.WithMessage(w, http.StatusCreated, msg)
}

// ViewProduct View all product
// @Summary View product
// @Description This endpoint for view all product.
// @Tags sale
// @Produce json
// @Success 200 {object} response.Base
// @Failure 400 {object} response.Base
// @Failure 404 {object} response.Base
// @Failure 500 {object} response.Base
// @Security BearerAuth
// @Router /v1/products [get]
func (h *SaleHandler) ViewProduct(w http.ResponseWriter, r *http.Request) {

	result, err := h.SaleService.GetProductList()
	if err != nil {
		response.WithError(w, err)
		return
	}
	response.WithJSON(w, http.StatusOK, result)
}
