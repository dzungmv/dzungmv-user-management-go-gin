package routes

import (
	"go/user-management/internal/handlers"
	"go/user-management/internal/middlewares"

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
	r.Use(
		middlewares.LoggerMiddleware(),
		middlewares.ApiKeyMiddleware(),
		middlewares.AuthMiddleware(),
		middlewares.RateLimiterMiddleware(),
	)

	users := r.Group("/users")
	{
		users.GET("", ur.handler.GetAllUsers)
		users.POST("", ur.handler.CreateUser)
		users.GET("/:uuid", ur.handler.GetUserByUuid)
	}
}
