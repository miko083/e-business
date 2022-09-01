package models

import (
	"github.com/jinzhu/gorm"
)

type ShippingCart struct {
	gorm.Model
	UserID               int                   `json:"user_id"`
	User                 User                  `json:"user"`
	ConsolesWithQuantity []ConsoleWithQuantity `json:"consoles_with_quantity"`
}
