package dto

import "go/user-management/internal/models"

type UserDTO struct {
	UUID   string `json:"uuid"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Age    int    `json:"age"`
	Status string `json:"status"`
	Level  string `json:"level"`
}

type CreateUserInput struct {
	UUID     string `json:"uuid"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email,email_advanced"`
	Age      int    `json:"age" binding:"required,gt=0"`
	Password string `json:"password" binding:"required,min=8,password_strong"`
	Status   int    `json:"status" binding:"required,oneof=1 2"`
	Level    int    `json:"level" binding:"required,oneof=1 2"`
}

func (u *CreateUserInput) MapCreateToUserModel() models.User {
	return models.User{
		Name:     u.Name,
		Email:    u.Email,
		Age:      u.Age,
		Password: u.Password,
		Status:   u.Status,
		Level:    u.Level,
	}
}

type UpdateUserInput struct {
	UUID     string `json:"uuid"`
	Name     string `json:"name" binding:"omitempty"`
	Email    string `json:"email" binding:"omitempty,email,email_advanced"`
	Age      int    `json:"age" binding:"omitempty,gt=0"`
	Password string `json:"password" binding:"omitempty,min=8,password_strong"`
	Status   int    `json:"status" binding:"omitempty,oneof=1 2"`
	Level    int    `json:"level" binding:"omitempty,oneof=1 2"`
}

func (u *UpdateUserInput) MapUpdateToUserModel() models.User {
	return models.User{
		Name:     u.Name,
		Email:    u.Email,
		Age:      u.Age,
		Password: u.Password,
		Status:   u.Status,
		Level:    u.Level,
	}
}

func MapUserToDTO(user models.User) *UserDTO {
	return &UserDTO{
		UUID:   user.UUID,
		Name:   user.Name,
		Email:  user.Email,
		Age:    user.Age,
		Status: mapStatusText(user.Status),
		Level:  mapLevelText(user.Level),
	}
}

func MapUsersToDTO(users []models.User) *[]UserDTO {
	usersDTO := make([]UserDTO, 0, len(users))

	for _, user := range users {
		usersDTO = append(usersDTO, *MapUserToDTO(user))
	}

	return &usersDTO
}

func mapStatusText(status int) string {
	switch status {
	case 1:
		return "Active"
	case 2:
		return "Inactive"
	default:
		return "None"
	}
}

func mapLevelText(level int) string {
	switch level {
	case 1:
		return "Admin"
	case 2:
		return "User"
	default:
		return "None"
	}
}
