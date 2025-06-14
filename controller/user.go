package controller

import (
	"net/http"
	"time"

	"github.com/Aman17101/Hospital/model"
	"github.com/Aman17101/Hospital/util"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (server Server) CreateUser(ctx *gin.Context) {
	util.Log(model.LogLevelInfo, model.ControllerPackage, model.CreateUser, " creating new user ", nil)
	user := model.User{}
	err := ctx.Bind(&user)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.CreateUser, "error while json binding ", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	user.ID = uuid.New()
	user.CreatedAt = time.Now()

	err = server.Postgressdb.CreateUser(&user)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.CreateUser, "error while inserting record", err)
		return
	}
	ctx.JSON(http.StatusCreated, user)

}
