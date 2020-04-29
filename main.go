package main

import (
	"MyFirstGinProject/config"
	"MyFirstGinProject/model"
	"MyFirstGinProject/router"
)

func main() {

	config.NewConfig()
	model.Setup()

	r := router.NewRouter()

	r.Run("localhost:8080")


	
}
