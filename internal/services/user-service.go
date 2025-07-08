package services

import (
	"go/user-management/internal/models"
	"go/user-management/internal/repositories"
	"go/user-management/internal/utils"
	"strings"

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

func (us *userService) GetAllUsers(search string, page, limit int) ([]models.User, error) {
	users, err := us.repo.FindAll()

	if err != nil {
		return []models.User{}, utils.NewError("failed to fetch users", string(utils.ErrCodeInternal))
	}

	var filterUsers []models.User

	if search != "" {
		search = strings.ToLower(search)

		for _, user := range users {
			name := strings.ToLower(user.Name)
			email := strings.ToLower(user.Email)

			if strings.Contains(name, search) || strings.Contains(email, search) {
				filterUsers = append(filterUsers, user)
			}
		}

	} else {
		filterUsers = users
	}

	start := (page - 1) * limit
	if start >= len(filterUsers) {
		return []models.User{}, nil
	}

	end := min(start+limit, len(filterUsers))

	return filterUsers[start:end], nil
}

func (us *userService) GetUserById(uuid string) (models.User, error) {
	user, found := us.repo.FindByUuid(uuid)

	if !found {
		return models.User{}, utils.NewError("user not found", string(utils.ErrCodeNotFound))
	}

	return user, nil
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
