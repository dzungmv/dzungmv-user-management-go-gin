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

func (ur *userRepository) FindByEmail(email string) (models.User, bool) {
	for _, user := range ur.users {
		if user.Email == email {
			return user, true
		}
	}

	return models.User{}, false
}

func (ur *userRepository) CreateUser(user models.User) error {
	ur.users = append(ur.users, user)

	return nil
}
