package repositories

import (
	"go/user-management/internal/models"
	"log"
)

type userRepository struct {
	users []models.User
}

func NewUserRepository() *userRepository {
	return &userRepository{
		users: make([]models.User, 0),
	}
}

func (ur *userRepository) FindAll() {
	log.Println("Find all users")
}
