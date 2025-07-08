package main

import (
	"go/user-management/internal/app"
	"go/user-management/internal/configs"
)

func main() {
	config := configs.NewConfig()

	application := app.NewApplication(config)

	if err := application.Run(); err != nil {
		panic(err)
	}
}
