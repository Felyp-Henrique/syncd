package services

import (
	"Felyp-Henrique/syncd/src/application/domain/entities"
	"Felyp-Henrique/syncd/src/application/domain/services"
	"Felyp-Henrique/syncd/src/application/features/repositories"
)

type UsersDataBaseService struct {
	repository *repositories.UsersDataBaseRepository
}

func NewUserDataBaseService(repository *repositories.UsersDataBaseRepository) services.UsersService {
	return &UsersDataBaseService{
		repository: repository,
	}
}

func (u *UsersDataBaseService) All() ([]*entities.User, error) {
	return u.repository.All()
}
