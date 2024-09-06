package services

import (
	"Felyp-Henrique/syncd/src/application/domains/entities"
	"Felyp-Henrique/syncd/src/application/domains/services"
	"Felyp-Henrique/syncd/src/application/features/repositories"
)

type UsersDataBaseService struct {
	repository *repositories.UsersDataBaseRepository
}

func NewUserDataBaseService(repository *repositories.UsersDataBaseRepository) services.IUsersService {
	return &UsersDataBaseService{
		repository: repository,
	}
}

func (u *UsersDataBaseService) All() ([]*entities.User, error) {
	return u.repository.All()
}
