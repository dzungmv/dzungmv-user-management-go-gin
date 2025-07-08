package app

import (
	"go/user-management/internal/handlers"
	"go/user-management/internal/repositories"
	"go/user-management/internal/routes"
	"go/user-management/internal/services"
)

type UserModule struct {
	routes routes.Route
}

func NewUserModule() *UserModule {

	userRepo := repositories.NewUserRepository()

	userService := services.NewUserService(userRepo)

	userHandler := handlers.NewUserHandler(userService)

	userRoutes := routes.NewUserRoutes(userHandler)

	return &UserModule{routes: userRoutes}
}

func (um *UserModule) Routes() routes.Route {
	return um.routes
}
