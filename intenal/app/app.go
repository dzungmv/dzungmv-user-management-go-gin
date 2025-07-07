package app

import (
	"go/user-management/intenal/configs"
	"go/user-management/intenal/routes"

	"github.com/gin-gonic/gin"
)

type Application struct {
	config *configs.Config
	router *gin.Engine
}

type Module interface {
	Routes() routes.Route
}

func NewApplication(config *configs.Config) *Application {
	r := gin.Default()

	modules := []Module{
		NewUserModule(),
	}

	routes.RegisterRoutes(r, GetModuleRoutes(modules)...)

	return &Application{
		config: config,
		router: r,
	}
}

func (a *Application) Run() error {
	return a.router.Run(a.config.ServerAddress)
}

func GetModuleRoutes(module []Module) []routes.Route {
	routeList := make([]routes.Route, len((module)))

	for i, module := range module {
		routeList[i] = module.Routes()
	}

	return routeList
}
