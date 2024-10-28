package services

import (
	"UnderProject/internal/core/models"
	"UnderProject/internal/core/ports"
)

type UserService struct {
	repo ports.UserInterface
}

func NewUserService(repo ports.UserInterface) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) UserCreate(user *models.User) error {
	return nil
}

func (s *UserService) UserGetById(id int) (*models.User, error) {
	return nil, nil
}

func (s *UserService) UserGetAllUser() ([]*models.User, error) {
	return nil, nil
}

func (s *UserService) UserUpdate(user *models.User) error {
	return nil
}

func (s *UserService) UserDelete(id int) error {
	return nil

}
