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

func (server Server) GetUsers(ctx *gin.Context) {
	util.Log(model.LogLevelInfo, model.ControllerPackage, model.GetUsers, " creating new user ", nil)

	users, err := server.Postgressdb.GetUsers()
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.GetUsers, "error while inserting record", err)
		return
	}
	ctx.JSON(http.StatusCreated, users)

}

func (server Server) GetUser(ctx *gin.Context) {
	util.Log(model.LogLevelInfo, model.ControllerPackage, model.GetUser, "fetching user by ID", nil)

	// // 1. Parse ID from URL
	id := ctx.Param("id")
	user, err := server.Postgressdb.GetUser(id)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.GetUser, "invalid user ID", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID"})
		return
	}

	// 3. Respond with user
	ctx.JSON(http.StatusOK, user)
}
