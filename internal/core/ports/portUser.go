package ports

import "UnderProject/internal/core/models"

type UserInterface interface {
	UserCreate(*models.User) error
	UserGetById(id int) (*models.User, error)
	UserGetAll() ([]*models.User, error)
	UserUpdate(*models.User) error
	UserDelete(int) error
}
