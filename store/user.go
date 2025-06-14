package store

import (
	"github.com/Aman17101/Hospital/model"
	"github.com/Aman17101/Hospital/util"
)

func (store Postgress) CreateUser(user *model.User) error {
	util.Log(model.LogLevelInfo, model.StorePackage, model.CreateUser, " creating new user  ", nil)
	response :=store.DB.Create(user)
	if response.Error != nil {
		util.Log(model.LogLevelError, model.StorePackage, model.CreateUser,"error while creating new user ",response.Error)
		return response.Error
	}

util.Log(model.LogLevelInfo, model.StorePackage, model.CreateUser,"created new user ",user )
return nil
}
