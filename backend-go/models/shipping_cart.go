package models

import (
	"github.com/jinzhu/gorm"
)

type ShippingCart struct {
	gorm.Model
	UserEmail            string                `json:"user_email"`
	User                 User                  `gorm:"foreignKey:UserEmail" json:"user"`
	ConsolesWithQuantity []ConsoleWithQuantity `json:"consoles_with_quantity"`
	PaymentDone          bool                  `json:"payment_done"`
}
