package repositories

import "go/user-management/internal/models"

type UserRepository interface {
	FindAll() ([]models.User, error)
	FindByEmail(email string) (models.User, bool)
	CreateUser(user models.User) error
	FindByUuid(uuid string) (models.User, bool)
	UpdateUser(user models.User) error
	DeleteUser(uuid string) error
}
