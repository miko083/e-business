package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	FirstName   string `json:"fist_name"`
	LastName    string `json:"last_name"`
	Email       string `gorm:"primaryKey" json:"email"`
	AccessToken string `json:"access_token"`
	LoginToken  string `json:"login_token"`
	IsLoggedIn  bool   `json:"if_logged_in"`
}
