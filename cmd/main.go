package main

import (
	"github.com/Aman17101/Hospital/api"
	"github.com/Aman17101/Hospital/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	api := api.ApiRoutes{}
	controller := controller.Server{}
	routes := gin.Default()
	api.StartApp(routes, controller)

	routes.Run(":8000") 
	//fmt.Printf("main server is %v \n", api)
}
