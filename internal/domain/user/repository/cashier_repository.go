package repository

import (
	"fmt"

	"github.com/mkp-pos-cashier-api/internal/domain/user/model"
	"github.com/mkp-pos-cashier-api/shared/failure"
	"github.com/rs/zerolog/log"
)

const (
	createUserQuery = `
	INSERT INTO "user" (id, username, password, role, created_at, created_by, updated_at, updated_by) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`
	checkUserbyUsernameQuery = `
	SELECT
		COUNT(username)
	FROM
		"user"
	WHERE
		username = $1
	`
	selectUserByUsername = `
	SELECT
		username,
		password,
		role
	FROM
		"user"
	WHERE
		username = $1
	`
)

type CashierRepository interface {
	CreateUser(createtUser *model.CreateUser) error
	GetUserByUsername(username string) (*model.UserByUsername, error)
}

func (r *UserRepositoryPostgres) CreateUser(createtUser *model.CreateUser) error {

	exist, err := r.IsExistUserByUsername(createtUser.Username)
	if err != nil {
		log.Error().Err(err).Msg("[CreateUser] Failed checking user whether already exists or not")
		return err
	}
	if exist {
		err = failure.Conflict("create", "user", "already exists")
		return err
	}

	query := fmt.Sprintf(createUserQuery)
	_, err = r.DB.Write.Exec(
		query,
		createtUser.Id,
		createtUser.Username,
		createtUser.Password,
		createtUser.Role,
		createtUser.CreatedAt,
		createtUser.CreatedBy,
		createtUser.UpdatedAt,
		createtUser.UpdatedBy,
	)
	if err != nil {
		log.Error().Err(err).Msg("[CreateUser] Failed exec create user query")
		return err
	}

	return nil
}

func (r *UserRepositoryPostgres) IsExistUserByUsername(username string) (bool, error) {
	query := fmt.Sprintf(checkUserbyUsernameQuery)
	count := 0
	err := r.DB.Read.Get(&count, query, username)
	if err != nil {
		log.Error().Err(err).Msg("[IsExistUserByUsername] Failed to check user")
		err = failure.InternalError(err)
		return false, err
	}
	return count > 0, nil
}

func (r *UserRepositoryPostgres) GetUserByUsername(username string) (*model.UserByUsername, error) {
	exist, err := r.IsExistUserByUsername(username)
	if err != nil {
		log.Error().Err(err).Msg("[GetUserByUsername] Failed checking user whether already exists or not")
		return nil, err
	}
	if !exist {
		err = failure.NotFound("User not found")
		return nil, err
	}

	query := fmt.Sprintf(selectUserByUsername)

	var user model.UserByUsername
	err = r.DB.Read.Get(&user, query, username)
	if err != nil {
		log.Error().Err(err).Msg("[GetUserByUsername] Failed to get user by username")
		err = failure.InternalError(err)
		return nil, err
	}
	return &user, nil
}
