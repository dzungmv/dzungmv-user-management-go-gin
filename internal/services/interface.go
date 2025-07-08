package services

import "go/user-management/internal/models"

type UserService interface {
	GetAllUsers() ([]models.User, error)
	CreateUser(user models.User) (models.User, error)
	GetUserById(uuid string) (models.User, error)
}
