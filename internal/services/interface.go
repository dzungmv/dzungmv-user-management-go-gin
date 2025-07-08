package services

import "go/user-management/internal/models"

type UserService interface {
	GetAllUsers()
	CreateUser(user models.User) (models.User, error)
}
