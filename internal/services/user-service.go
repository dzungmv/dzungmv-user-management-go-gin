package services

import (
	"go/user-management/internal/models"
	"go/user-management/internal/repositories"
	"go/user-management/internal/utils"
	"log"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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

func (us *userService) CreateUser(user models.User) (models.User, error) {
	user.Email = utils.NormalizeString(user.Email)

	if _, exist := us.repo.FindByEmail(user.Email); exist {
		return models.User{}, utils.NewError("email aldready exits", string(utils.ErrCodeConflict))
	}

	user.UUID = uuid.New().String()

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return models.User{}, utils.WrapError(err, "failed to create user", string(utils.ErrCodeInternal))
	}

	user.Password = string(hashPassword)

	if err := us.repo.CreateUser(user); err != nil {
		return models.User{}, utils.WrapError(err, "failed to create user", string(utils.ErrCodeInternal))
	}

	return user, nil
}
