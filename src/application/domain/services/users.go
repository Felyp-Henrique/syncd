package services

import "Felyp-Henrique/syncd/src/application/domain/entities"

type UsersService interface {
	All() ([]*entities.User, error)
}
