package services

import "Felyp-Henrique/syncd/src/application/domains/entities"

type IUsersService interface {
	All() ([]*entities.User, error)
}
