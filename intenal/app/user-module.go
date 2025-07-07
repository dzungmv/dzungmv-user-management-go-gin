package app

import (
	"go/user-management/intenal/handlers"
	"go/user-management/intenal/repositories"
	"go/user-management/intenal/routes"
	"go/user-management/intenal/services"
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
