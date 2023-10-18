package infras

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/mkp-pos-cashier-api/configs"
	"github.com/rs/zerolog/log"
)

const (
	maxIdleConnection = 10
	maxOpenConnection = 10
)

// PostgreSQLConn wraps a pair of read/write MySQL connections.
type PostgreSQLConn struct {
	Read  *sqlx.DB
	Write *sqlx.DB
}

// ProvidePostgreSQLConn is the provider for PostgreSQLConn.
func ProvidePostgreSQLConn(config *configs.Config) *PostgreSQLConn {
	return &PostgreSQLConn{
		Read:  CreatePostgreSQLReadConn(*config),
		Write: CreatePostgreSQLWriteConn(*config),
	}
}

// CreatePostgreSQLWriteConn creates a database connection for write access.
func CreatePostgreSQLWriteConn(config configs.Config) *sqlx.DB {
	return CreateDBConnection(
		"write",
		config.DB.PostgreSQL.Write.Username,
		config.DB.PostgreSQL.Write.Password,
		config.DB.PostgreSQL.Write.Host,
		config.DB.PostgreSQL.Write.Port,
		config.DB.PostgreSQL.Write.Name,
		config.DB.PostgreSQL.Write.Timezone,
		"disable",
	)
}

// CreatePostgreSQLReadConn creates a database connection for read access.
func CreatePostgreSQLReadConn(config configs.Config) *sqlx.DB {
	return CreateDBConnection(
		"read",
		config.DB.PostgreSQL.Read.Username,
		config.DB.PostgreSQL.Read.Password,
		config.DB.PostgreSQL.Read.Host,
		config.DB.PostgreSQL.Read.Port,
		config.DB.PostgreSQL.Read.Name,
		config.DB.PostgreSQL.Read.Timezone,
		"disable",
	)
}

// CreateDBConnection creates a database connection.
func CreateDBConnection(name, username, password, host, port, dbName, timeZone, sslMode string) *sqlx.DB {
	descriptor := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host,
		port,
		username,
		password,
		dbName,
		sslMode,
	)
	db, err := sqlx.Connect("postgres", descriptor)
	if err != nil {
		log.
			Fatal().
			Err(err).
			Str("name", name).
			Str("host", host).
			Str("port", port).
			Str("dbName", dbName).
			Msg("Failed connecting to database")
	} else {
		log.
			Info().
			Str("name", name).
			Str("host", host).
			Str("port", port).
			Str("dbName", dbName).
			Msg("Connected to database")
	}
	db.SetMaxIdleConns(maxIdleConnection)
	db.SetMaxOpenConns(maxOpenConnection)

	return db
}
