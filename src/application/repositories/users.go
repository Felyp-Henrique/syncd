package repositories

import (
	"Felyp-Henrique/syncd/src/application/domains/entities"
	"database/sql"
)

type UsersDataBaseRepository struct {
	database *sql.DB
}

func NewUsersDataBaseRepository(database *sql.DB) *UsersDataBaseRepository {
	return &UsersDataBaseRepository{
		database: database,
	}
}

func (u *UsersDataBaseRepository) All() ([]entities.User, error) {
	return []entities.User{}, nil
}
