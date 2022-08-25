package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	FirstName string `json:"fist_name"`
	LastName  string `json:"last_name"`
}
