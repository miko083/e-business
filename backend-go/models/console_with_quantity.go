package models

import (
	"github.com/jinzhu/gorm"
)

type ConsoleWithQuantity struct {
	gorm.Model
	ConsoleID      int     `json:"console_id"`
	Console        Console `json:"console"`
	Quantity       int     `json:"quantity"`
	ShippingCartID uint    `json:"shipping_cart_id"`
}
