package model

import (
	"time"

	"github.com/gofrs/uuid"
)

type CreateUser struct {
	Id        uuid.UUID `db:"id"`
	Username  string    `db:"username"`
	Password  string    `db:"password"`
	Role      string    `db:"role"`
	CreatedAt time.Time `db:"created_at"`
	CreatedBy string    `db:"Created_by"`
	UpdatedAt time.Time `db:"updated_at"`
	UpdatedBy string    `db:"updated_by"`
}

type CreateUserList []*CreateUser

type UserByUsername struct {
	Username string `db:"username"`
	Password string `db:"password"`
	Role     string `db:"role"`
}

type UserByUsernameList []*UserByUsername
