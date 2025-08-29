package store

import (
	"github.com/Aman17101/Hospital/model"
	"github.com/Aman17101/Hospital/util"
)

func (store Postgress) CreateUser(user *model.User) error {
	util.Log(model.LogLevelInfo, model.StorePackage, model.CreateUser, " creating new user  ", nil)
	response := store.DB.Create(user)
	if response.Error != nil {
		util.Log(model.LogLevelError, model.StorePackage, model.NewStore, "error while creating new user ", response.Error)
		return response.Error
	}

	util.Log(model.LogLevelInfo, model.StorePackage, model.CreateUser, "created new user ", user)
	return nil
}

func (store Postgress) GetUsers() ([]model.User, error) {
	users := []model.User{}

	util.Log(model.LogLevelInfo, model.StorePackage, model.GetUsers, " fetching records of user from db ", nil)
	if err := store.DB.Find(&users).Error; err != nil {
		util.Log(model.LogLevelError, model.StorePackage, model.GetUsers, "error while fetching user from db ", err)
		return users, err
	}

	util.Log(model.LogLevelInfo, model.StorePackage, model.GetUsers, "creatednew user ", users)
	return users, nil
}

func (store Postgress) GetUser(id string) (model.User, error) {
	var user model.User

	util.Log(model.LogLevelInfo, model.StorePackage, model.GetUser, "fetching user by ID from db", map[string]interface{}{"id": id})

	if err := store.DB.Where("id = ?", id).First(&user).Error; err != nil {
		util.Log(model.LogLevelError, model.StorePackage, model.GetUser, "error while fetching user by ID from db", err)
		return user, err
	}

	util.Log(model.LogLevelInfo, model.StorePackage, model.GetUser, "user fetched successfully", user)
	return user, nil
}
