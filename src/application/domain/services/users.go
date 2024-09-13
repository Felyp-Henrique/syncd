package services

import "Felyp-Henrique/syncd/src/application/domain/entities"

type IUsersService interface {
	All() ([]*entities.User, error)
}
