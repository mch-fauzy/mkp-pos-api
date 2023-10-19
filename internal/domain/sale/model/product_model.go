package model

import (
	"time"

	"github.com/guregu/null"
)

type CreateProduct struct {
	Id        int       `db:"id"`
	Name      string    `db:"name"`
	Category  string    `db:"category"`
	Stock     int       `db:"stock"`
	CreatedAt time.Time `db:"created_at"`
	CreatedBy string    `db:"created_by"`
	UpdatedAt time.Time `db:"updated_at"`
	UpdatedBy string    `db:"updated_by"`
}

type Product struct {
	Id        int         `db:"id"`
	Name      string      `db:"name"`
	Category  string      `db:"category"`
	Stock     int         `db:"stock"`
	CreatedAt time.Time   `db:"created_at"`
	CreatedBy string      `db:"created_by"`
	UpdatedAt time.Time   `db:"updated_at"`
	UpdatedBy string      `db:"updated_by"`
	DeletedAt null.Time   `db:"deleted_at"`
	DeletedBy null.String `db:"deleted_by"`
}

type ProductList []*Product
