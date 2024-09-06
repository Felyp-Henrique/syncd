package repositories

import (
	"Felyp-Henrique/syncd/src/application/domains/entities"
	"strings"

	"github.com/jmoiron/sqlx"
)

type UsersDataBaseRepository struct {
	database *sqlx.DB
}

func NewUsersDataBaseRepository(database *sqlx.DB) *UsersDataBaseRepository {
	return &UsersDataBaseRepository{
		database: database,
	}
}

func (u *UsersDataBaseRepository) All() ([]*entities.User, error) {
	var (
		users []*entities.User = []*entities.User{}
		user  *entities.User   = nil
		query *strings.Builder = new(strings.Builder)
		rows  *sqlx.Rows       = nil
		err   error            = nil
	)
	query.WriteString("select\n")
	query.WriteString("  u.id,\n")
	query.WriteString("  u.name,\n")
	query.WriteString("  u.email,\n")
	query.WriteString("  u.password,\n")
	query.WriteString("  u.type,\n")
	query.WriteString("  u.status\n")
	query.WriteString("from\n")
	query.WriteString("	 users u\n")
	query.WriteString("where\n")
	query.WriteString("	 u.status = :status\n")
	rows, err = u.database.NamedQuery(query.String(), map[string]any{
		"status": entities.USER_STATUS_ACTIVED,
	})
	if err != nil {
		return []*entities.User{}, err
	}
	for rows.Next() {
		user = new(entities.User)
		rows.StructScan(user)
		users = append(users, user)
	}
	return users, nil
}
