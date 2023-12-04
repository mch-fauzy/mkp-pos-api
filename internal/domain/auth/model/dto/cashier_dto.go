package dto

import (
	"time"

	"github.com/gofrs/uuid"
	"github.com/mkp-pos-cashier-api/internal/domain/auth/model"
	"github.com/mkp-pos-cashier-api/shared"
	"github.com/mkp-pos-cashier-api/shared/failure"
)

type RegisterCashierRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (r RegisterCashierRequest) Validate() error {
	if r.Username == "" {
		return failure.BadRequestFromString("Username is required")
	}

	if r.Password == "" {
		return failure.BadRequestFromString("Password is required")
	}

	if len(r.Password) < 8 {
		return failure.BadRequestFromString("Password must be at least 8 characters")
	}

	return nil
}

func (r RegisterCashierRequest) ToModel() model.CreateUser {
	id, _ := uuid.NewV4()
	currentTime := time.Now()
	username := r.Username
	return model.CreateUser{
		Id:        id,
		Username:  username,
		Password:  r.Password,
		Role:      shared.CashierRole,
		CreatedAt: currentTime,
		CreatedBy: username,
		UpdatedAt: currentTime,
		UpdatedBy: username,
	}
}

type LoginCashierRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (r LoginCashierRequest) Validate() error {
	if r.Username == "" {
		return failure.BadRequestFromString("Username is required")
	}

	if r.Password == "" {
		return failure.BadRequestFromString("Password is required")
	}

	return nil
}

type LoginResponse struct {
	Token string `json:"token"`
}

func BuildLoginResponse(token string) LoginResponse {
	return LoginResponse{
		Token: token,
	}
}
