package api

import (
	"github.com/Aman17101/Hospital/controller"
	"github.com/Aman17101/Hospital/store"
	"github.com/gin-gonic/gin"
)

type ApiRoutes struct {
	Server controller.ServerOperations
}

func (api *ApiRoutes) StartApp(routes *gin.Engine, server controller.Server) {

	api.Server = &server
	api.Server.NewServer(store.Postgress{})

	api.UserRouts(routes)

}
