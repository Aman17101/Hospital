package api

import (
	"github.com/Aman17101/Hospital/controller"
	"github.com/Aman17101/Hospital/store"
)

type ApiRoutes struct {
	Server controller.ServerOperations
}

func (api *ApiRoutes)StartApp(server controller.Server){

	api.Server=&server
	api.Server.NewServer(store.Postgress{})

}