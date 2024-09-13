package database

import "github.com/jmoiron/sqlx"

type DataBase interface {
	GetConnection() (*sqlx.DB, error)
}
