package models

import (
	"github.com/jinzhu/gorm"
)

type Console struct {
	gorm.Model
	Name           string       `json:"name"`
	Price          float64      `json:"price"`
	ManufacturerID int          `json:"manufacturer_id"`
	Manufacturer   Manufacturer `json:"manufacturer"`
}
