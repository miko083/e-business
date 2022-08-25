package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	DBconnection *gorm.DB
)

func Connect() {
	var err error
	DBconnection, err = gorm.Open("sqlite3", "console_shop.db")
	if err != nil {
		panic("Can't connect to database!")
	}
	DBconnection.LogMode(true)
}
