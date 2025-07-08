package repositories

import (
	"fmt"
	"go/user-management/internal/models"
)

type userRepository struct {
	users []models.User
}

func NewUserRepository() *userRepository {
	return &userRepository{
		users: make([]models.User, 0),
	}
}

func (ur *userRepository) FindAll() ([]models.User, error) {
	return ur.users, nil
}

func (ur *userRepository) FindByUuid(uuid string) (models.User, bool) {
	for _, user := range ur.users {
		if user.UUID == uuid {
			return user, true
		}
	}
	return models.User{}, false
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

func (ur *userRepository) UpdateUser(user models.User) error {
	for index, _user := range ur.users {
		if _user.UUID == user.UUID {
			ur.users[index] = user
			return nil
		}
	}

	return fmt.Errorf("user not found")
}
