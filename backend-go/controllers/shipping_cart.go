package controllers

import (
	m "consoleshop/models"
	"log"
	"net/http"

	"consoleshop/database"

	"github.com/labstack/echo/v4"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func GetCarts(c echo.Context) error {
	var shippingCartsWithQuantity []m.ShippingCart
	database.DBconnection.Preload("User").Preload("ConsolesWithQuantity.Console").Preload("ConsolesWithQuantity.Console.Manufacturer").Find(&shippingCartsWithQuantity)
	return c.JSON(http.StatusOK, shippingCartsWithQuantity)
}

func GetCart(c echo.Context) error {
	id := c.Param("id")
	var shippingCart m.ShippingCart
	database.DBconnection.Preload("User").Preload("ConsolesWithQuantity").Preload("ConsolesWithQuantity.Console").Preload("ConsolesWithQuantity.Console.Manufacturer").Find(&shippingCart, "ID = ?", id)
	return c.JSON(http.StatusOK, shippingCart)
}

func AddCart(c echo.Context) error {
	shippingCart := m.ShippingCart{}
	err := c.Bind(&shippingCart)
	if err != nil {
		log.Printf("Failed: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	database.DBconnection.Create(&shippingCart)
	return c.JSON(http.StatusOK, "Added new shipping cart.")
}

func DeleteCart(c echo.Context) error {
	id := c.Param("id")
	database.DBconnection.Delete(&m.ShippingCart{}, id)
	return c.JSON(http.StatusOK, "Deleted shipping cart with the id: "+id)
}

func UpdateCart(c echo.Context) error {
	id := c.Param("id")
	var shippingCartToUpdate m.ShippingCart
	database.DBconnection.Find(&shippingCartToUpdate, "ID = ?", id)

	shippingCartFromBody := m.ShippingCart{}
	err := c.Bind(&shippingCartFromBody)
	if err != nil {
		log.Printf("Failed: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	if shippingCartFromBody.UserID != 0 {
		shippingCartToUpdate.UserID = shippingCartFromBody.UserID
	}

	if shippingCartFromBody.ConsolesWithQuantity != nil {
		shippingCartToUpdate.ConsolesWithQuantity = shippingCartFromBody.ConsolesWithQuantity
	}

	database.DBconnection.Save(&shippingCartToUpdate)
	return c.JSON(http.StatusOK, "Updated shipping cart with the id: "+id)
}
