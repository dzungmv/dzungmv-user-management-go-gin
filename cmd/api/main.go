package main

import (
	"go/user-management/intenal/app"
	"go/user-management/intenal/configs"
)

func main() {
	config := configs.NewConfig()

	application := app.NewApplication(config)

	if err := application.Run(); err != nil {
		panic(err)
	}
}
