package handlers

import (
	"go/user-management/internal/models"
	"go/user-management/internal/services"
	"go/user-management/internal/utils"
	"go/user-management/internal/validations"
	"net/http"

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

}

func (uh *UserHandler) CreateUser(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadGateway, validations.HandleValidationErrors(err))
	}

	user, err := uh.service.CreateUser(user)

	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	utils.ResponseSuccess(ctx, http.StatusCreated, user)

}
