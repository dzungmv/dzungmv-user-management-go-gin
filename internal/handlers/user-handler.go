package handlers

import (
	"go/user-management/internal/dto"
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

type GetUsersParams struct {
	Search string `form:"search" binding:"omitempty,min=3,max=50,search"`
	Page   int    `form:"page" binding:"omitempty,gte=1,lte=100"`
	Limit  int    `form:"limit" binding:"omitempty,gte=1,lte=100"`
}

func NewUserHandler(service services.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (uh *UserHandler) GetAllUsers(ctx *gin.Context) {
	var params GetUsersParams

	if err := ctx.ShouldBindQuery(&params); err != nil {
		utils.ResponseValidator(ctx, validations.HandleValidationErrors(err))
		return
	}

	if params.Page == 0 {
		params.Page = 1
	}

	if params.Limit == 0 {
		params.Limit = 10
	}

	users, err := uh.service.GetAllUsers(params.Search, params.Page, params.Limit)

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
	var input dto.CreateUserInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		utils.ResponseValidator(ctx, validations.HandleValidationErrors(err))
		return
	}

	user := input.MapCreateToUserModel()

	created, err := uh.service.CreateUser(user)

	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	userDTO := dto.MapUserToDTO(created)

	utils.ResponseSuccess(ctx, http.StatusCreated, &userDTO)

}

func (uh *UserHandler) UpdateUser(ctx *gin.Context) {
	var params GetUserByUuidParam
	if err := ctx.ShouldBindUri(&params); err != nil {
		utils.ResponseValidator(ctx, validations.HandleValidationErrors(err))
		return
	}

	var input dto.UpdateUserInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		utils.ResponseValidator(ctx, validations.HandleValidationErrors(err))
		return
	}

	user := input.MapUpdateToUserModel()

	updatedUser, err := uh.service.UpdateUser(params.UUID, user)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	userDTO := dto.MapUserToDTO(updatedUser)

	utils.ResponseSuccess(ctx, http.StatusOK, userDTO)

}

func (uh *UserHandler) DeleteUser(ctx *gin.Context) {

	var params GetUserByUuidParam
	if err := ctx.ShouldBindUri(&params); err != nil {
		utils.ResponseValidator(ctx, validations.HandleValidationErrors(err))
		return
	}

	if err := uh.service.DeleteUser(params.UUID); err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	utils.ResponseSuccess(ctx, http.StatusFound, gin.H{
		"message": "detele user successfully",
	})
}
