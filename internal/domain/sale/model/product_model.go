package model

import (
	"time"
)

type CreateProduct struct {
	Id        int       `db:"id"`
	Name      string    `db:"name"`
	Category  string    `db:"category"`
	Stock     int       `db:"stock"`
	CreatedAt time.Time `db:"created_at"`
	CreatedBy string    `db:"Created_by"`
	UpdatedAt time.Time `db:"updated_at"`
	UpdatedBy string    `db:"updated_by"`
}
