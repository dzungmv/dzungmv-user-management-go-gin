package handlers

import (
	"go/user-management/internal/dto"
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

type GetUserByUuidParam struct {
	UUID string `uri:"uuid" binding:"uuid"`
}

func NewUserHandler(service services.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (uh *UserHandler) GetAllUsers(ctx *gin.Context) {
	users, err := uh.service.GetAllUsers()

	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	usersDTO := dto.MapUsersToDTO(users)

	utils.ResponseSuccess(ctx, http.StatusOK, &usersDTO)
}

func (uh *UserHandler) GetUserByUuid(ctx *gin.Context) {
	var params GetUserByUuidParam

	if err := ctx.ShouldBindUri(&params); err != nil {
		utils.ResponseValidator(ctx, validations.HandleValidationErrors(err))
		return
	}

	user, err := uh.service.GetUserById(params.UUID)

	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	userDTO := dto.MapUserToDTO(user)

	utils.ResponseSuccess(ctx, http.StatusOK, &userDTO)
}

func (uh *UserHandler) CreateUser(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		utils.ResponseValidator(ctx, validations.HandleValidationErrors(err))
		return
	}

	user, err := uh.service.CreateUser(user)

	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	userDTO := dto.MapUserToDTO(user)

	utils.ResponseSuccess(ctx, http.StatusCreated, &userDTO)

}
