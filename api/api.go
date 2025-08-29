package api

import (
	"github.com/Aman17101/Hospital/controller"
	"github.com/Aman17101/Hospital/store"
	"github.com/gin-gonic/gin"

	_ "github.com/Aman17101/Hospital/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type ApiRoutes struct {
	Server controller.ServerOperations
}

func (api *ApiRoutes) StartApp(routes *gin.Engine, server controller.Server) {

	api.Server = &server
	api.Server.NewServer(store.Postgress{})

	//swagger documentation
	routes.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api.UserRouts(routes)

}
