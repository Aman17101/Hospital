package main

import (
	"github.com/Aman17101/Hospital/api"
	"github.com/Aman17101/Hospital/controller"
	"github.com/gin-gonic/gin"
)
// @title testdb
// @description API for managing Hospital
// @host localhost:8000
// @BasePath /
// @schemes http https

func main() {
	api := api.ApiRoutes{}
	controller := controller.Server{}
	routes := gin.Default()
	api.StartApp(routes, controller)

	routes.Run(":8000") 
	//fmt.Printf("main server is %v \n", api)
}
