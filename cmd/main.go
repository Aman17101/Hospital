package main

import (
	"fmt"

	"github.com/Aman17101/Hospital/api"
	"github.com/Aman17101/Hospital/controller"
)

func main() {
	api := api.ApiRoutes{}
	api.StartApp(controller.Server{})
	fmt.Printf("main server is %v \n", api)
}
