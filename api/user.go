package api

import (
	"github.com/Aman17101/Hospital/model"
	"github.com/Aman17101/Hospital/util"
	"github.com/gin-gonic/gin"
)

func (api ApiRoutes) UserRouts(routes *gin.Engine) {
	group := routes.Group("user")
	{
		group.POST("/create", api.CreateUser)
	}

}
func (api ApiRoutes) CreateUser(ctx *gin.Context) {
	util.Log(model.LogLevelInfo, model.ApiPackage, model.CreateUser, "creating new user", nil)
	api.Server.CreateUser(ctx)
}
