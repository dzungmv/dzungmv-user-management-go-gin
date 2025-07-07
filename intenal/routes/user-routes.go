package routes

import (
	"go/user-management/intenal/handlers"
	"go/user-management/intenal/middlewares"

	"github.com/gin-gonic/gin"
)

type UserRoutes struct {
	handler *handlers.UserHandler
}

func NewUserRoutes(handler *handlers.UserHandler) *UserRoutes {
	return &UserRoutes{
		handler: handler,
	}
}

func (ur *UserRoutes) Register(r *gin.RouterGroup) {
	r.Use(middlewares.AuthMiddleware())
	users := r.Group("/users")
	{
		users.GET("", ur.handler.GetAllUsers)
	}
}
