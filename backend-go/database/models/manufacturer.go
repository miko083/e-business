package models

import (
	"github.com/jinzhu/gorm"
)

type Manufacturer struct {
	gorm.Model
	Name          string `json:"name"`
	OriginCountry string `json:"origin_country"`
}
