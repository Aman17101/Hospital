package controller

import (
	"fmt"

	"github.com/Aman17101/Hospital/model"
	"github.com/Aman17101/Hospital/store"
	"github.com/Aman17101/Hospital/util"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Postgressdb store.StoreOperation
}

func (s *Server) NewServer(pgstore store.Postgress) {

	util.SetLogger()
	util.Logger.Infof("Logger setup done \n")
	s.Postgressdb = &pgstore
	err := s.Postgressdb.NewStore()
	if err != nil {
		util.Logger.Errorf("Error while creating store connection ,err %v\n", err)
		util.Log(model.LogLevelError, model.Controller, "NewServer", "Error while creating store connection \n", err)
	} else {
		util.Logger.Infof("Connected with store \n")
		util.Log(model.LogLevelInfo, model.Controller, model.NewServer, "Connected with store \n", nil)
	}

	fmt.Printf("server is %v \n", s)
}

type ServerOperations interface {
	NewServer(pgstore store.Postgress)
	CreateUser(ctx *gin.Context)
}
