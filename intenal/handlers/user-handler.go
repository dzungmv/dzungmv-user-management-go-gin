package handlers

import (
	"go/user-management/intenal/services"
	"log"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service services.UserService
}

func NewUserHandler(service services.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (uh *UserHandler) GetAllUsers(ctx *gin.Context) {
	log.Println("Get all users handlers")

	uh.service.GetAllUsers()
}
