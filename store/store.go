package store

import (
	"fmt"

	"github.com/Aman17101/Hospital/model"
	"github.com/Aman17101/Hospital/util"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgress struct {
	DB *gorm.DB
}

func (store *Postgress) NewStore() error {
	dsn := "host=localhost user=vishal password=password dbname=manage port=5432 sslmode=disable"
	util.Log(model.LogLevelInfo, model.StorePackage, model.NewStore, "creating new store ", nil)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		util.Log(model.LogLevelError, model.StorePackage, model.NewStore, "error while creating store ", err)
		return err
	} else {
		store.DB = db
	}
	err = db.AutoMigrate(
		model.User{},
	)
	if err != nil {
		util.Log(model.LogLevelError, model.StorePackage, model.NewStore, "error while creating automigration", err)
		return err
	}
	fmt.Printf("db is %v \n", db)
	return nil

}

type StoreOperation interface {
	NewStore() error
	CreateUser(user *model.User) error
	GetUser(id string)(model.User,error)
	GetUsers()([]model.User,error)
}
