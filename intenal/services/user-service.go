package services

import (
	"go/user-management/intenal/repositories"
	"log"
)

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

func (us *userService) GetAllUsers() {
	log.Println("Get all users services")

	us.repo.FindAll()
}
