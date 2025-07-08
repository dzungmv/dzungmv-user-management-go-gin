package app

import (
	"go/user-management/internal/configs"
	"go/user-management/internal/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Application struct {
	config  *configs.Config
	router  *gin.Engine
	modules []Module
}

type Module interface {
	Routes() routes.Route
}

func NewApplication(config *configs.Config) *Application {
	r := gin.Default()

	loadEnv()

	modules := []Module{
		NewUserModule(),
	}

	routes.RegisterRoutes(r, GetModuleRoutes(modules)...)

	return &Application{
		config:  config,
		router:  r,
		modules: modules,
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

func loadEnv() {
	err := godotenv.Load("../../.env")

	if err != nil {
		log.Println("Can not find .env file, try again!")
	}

}
