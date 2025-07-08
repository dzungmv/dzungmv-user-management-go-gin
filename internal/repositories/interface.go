package repositories

import "go/user-management/internal/models"

type UserRepository interface {
	FindAll()
	FindByEmail(email string) (models.User, bool)
	CreateUser(user models.User) error
}
