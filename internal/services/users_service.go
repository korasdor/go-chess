package services

import (
	"github.com/korasdor/go-chess/internal/domain"
	"github.com/korasdor/go-chess/internal/repository"
)

type UsersService struct {
	repo *repository.Repositories
}

func NewUsersService(repo *repository.Repositories) *UsersService {
	return &UsersService{
		repo: repo,
	}
}

func (s *UsersService) GetUser(userId int) (domain.UserData, error) {
	return s.repo.UsersRepo.GetById(userId)
}

func (s *UsersService) UpdateUser(user domain.UserData) (domain.UserData, error) {
	return s.repo.UsersRepo.UpdateUser(user)
}

func (s *UsersService) DeleteUser(userId int) error {
	return nil
}
