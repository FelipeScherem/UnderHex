package mysql

import (
	"UnderProject/internal/core/models"
	"UnderProject/internal/core/ports"
)

type UserInterface struct {
}

func NewUserRepositoryMySql() ports.UserInterface {
	return &UserInterface{}
}

func (u UserInterface) UserCreate(user *models.User) error {
	return nil
}

func (u UserInterface) UserGetById(id int) (*models.User, error) {
	return nil, nil
}

func (u UserInterface) UserGetAll() ([]*models.User, error) {
	return nil, nil
}

func (u UserInterface) UserUpdate(user *models.User) error {
	return nil
}

func (u UserInterface) UserDelete(int) error {
	return nil
}
