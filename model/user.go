package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"primaryKey" json:"id"`
	Name               `gorm:"embedded" json:"name"`
	Address   Address  `gorm:"embedded" json:"address"`
	CreatedBy string   `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedBy string   `json:"updated_by"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedBy string   `json:"deleted_by"`
	DeletedAt time.Time `json:"deleted_at"`
}

type Name struct {
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	LastName   string `json:"last_name"`
}

type Address struct {
	Lane     string `json:"lane"`
	Village  string `json:"village"`
	City     string `json:"city"`
	District string `json:"district"`
	Pincode  int    `json:"pincode"`
	State    string `json:"state"`
}
